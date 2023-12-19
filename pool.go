package main

import (
	"context"
	"fmt"
	"sync"
)

// WorkerManager manages the workers
type WorkerPool struct {
	waitGroup         sync.WaitGroup
	context           context.Context
	contextCancelFunc context.CancelFunc
	observer          *WorkerExecutionObserver
	workers           []Worker
}

func NewWorkerPool() *WorkerPool {
	return &WorkerPool{
		observer: NewWorkerExecutionObserver(),
		workers:  []Worker{},
	}
}

func (p *WorkerPool) Observer() *WorkerExecutionObserver {
	return p.observer
}

func (p *WorkerPool) Start(ctx context.Context) error {
	p.context, p.contextCancelFunc = context.WithCancel(ctx)

	for _, worker := range p.workers {
		p.startWorker(p.context, worker)
	}

	return nil
}

func (p *WorkerPool) Stop() error {
	p.contextCancelFunc()

	p.waitGroup.Wait()

	return nil
}

func (p *WorkerPool) AddWorkers(workers ...Worker) {
	p.workers = append(p.workers, workers...)
}

func (p *WorkerPool) startWorker(ctx context.Context, worker Worker) {
	p.waitGroup.Add(1)

	go func() {
		defer func() {
			p.waitGroup.Done()

			if r := recover(); r != nil {
				p.observer.Observe(worker, Stopped, fmt.Sprintf("panic recovered: %s", r))
			}
		}()

		p.observer.Observe(worker, Running, "started")

		if err := worker.Run(ctx); err != nil {
			p.observer.Observe(worker, Stopped, err.Error())
		} else {
			p.observer.Observe(worker, Stopped, "completed")
		}
	}()
}

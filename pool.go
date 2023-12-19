package main

import (
	"context"
	"fmt"
	"sync"
)

// WorkerPool workers pool
type WorkerPool struct {
	mutex             sync.Mutex
	waitGroup         sync.WaitGroup
	context           context.Context
	contextCancelFunc context.CancelFunc
	workers           []Worker
	executions        map[string]*WorkerExecution
	executionsLimit   int
}

func NewWorkerPool(options ...WorkerPoolOption) *WorkerPool {
	appliedOptions := DefaultWorkerPoolOptions()
	for _, opt := range options {
		opt(&appliedOptions)
	}

	return &WorkerPool{
		workers:         []Worker{},
		executions:      make(map[string]*WorkerExecution),
		executionsLimit: appliedOptions.ExecutionsLimit,
	}
}

func (p *WorkerPool) AddWorkers(workers ...Worker) *WorkerPool {
	p.workers = append(p.workers, workers...)

	return p
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

func (p *WorkerPool) Report() map[string]*WorkerExecution {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	return p.executions
}

func (p *WorkerPool) startWorker(ctx context.Context, worker Worker) {
	p.waitGroup.Add(1)

	workerExecution := p.workerExecution(worker)

	go func(ctx context.Context, workerExecution *WorkerExecution) {
		defer func() {
			p.waitGroup.Done()

			if r := recover(); r != nil {
				workerExecution.
					SetStatus(Stopped).
					AddEvent(fmt.Sprintf("stopping execution %d/%d with recovered panic: %s", workerExecution.Current(), workerExecution.Limit(), r))

				if workerExecution.Current() < workerExecution.Limit() {
					workerExecution.AddEvent("restarting after panic recovery")

					p.startWorker(ctx, worker)
				}
			}
		}()

		workerExecution.
			SetStatus(Running).
			SetCurrent(workerExecution.Current() + 1).
			AddEvent(fmt.Sprintf("starting execution %d/%d", workerExecution.Current(), workerExecution.Limit()))

		if err := worker.Run(ctx); err != nil {
			workerExecution.
				SetStatus(Stopped).
				AddEvent(fmt.Sprintf("stopping execution %d/%d with error: %v", workerExecution.Current(), workerExecution.Limit(), err.Error()))

			if workerExecution.Current() < workerExecution.Limit() {
				workerExecution.AddEvent("restarting after error")

				p.startWorker(ctx, worker)
			}
		} else {
			workerExecution.
				SetStatus(Stopped).
				AddEvent(fmt.Sprintf("stopping execution %d/%d with success", workerExecution.Current(), workerExecution.Limit()))
		}
	}(ctx, workerExecution)
}

func (p *WorkerPool) workerExecution(worker Worker) *WorkerExecution {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	if _, ok := p.executions[worker.Name()]; !ok {
		p.executions[worker.Name()] = NewWorkerExecution(worker, p.executionsLimit)
	}

	return p.executions[worker.Name()]
}

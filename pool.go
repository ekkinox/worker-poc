package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// WorkerPool workers pool
type WorkerPool struct {
	mutex             sync.Mutex
	waitGroup         sync.WaitGroup
	context           context.Context
	contextCancelFunc context.CancelFunc
	options           Options
	workers           []*ResolvedWorker
	executions        map[string]*WorkerExecution
}

func NewWorkerPool(options ...WorkerOption) *WorkerPool {
	globalOptions := DefaultWorkerOptions()
	for _, opt := range options {
		opt(&globalOptions)
	}

	return &WorkerPool{
		workers:    []*ResolvedWorker{},
		executions: make(map[string]*WorkerExecution),
		options:    globalOptions,
	}
}

func (p *WorkerPool) AddResolvedWorkers(workers ...*ResolvedWorker) *WorkerPool {
	p.workers = append(p.workers, workers...)

	return p
}

func (p *WorkerPool) Start(ctx context.Context) error {
	p.context, p.contextCancelFunc = context.WithCancel(ctx)

	for _, worker := range p.workers {
		p.startResolvedWorker(p.context, worker)
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

func (p *WorkerPool) startResolvedWorker(ctx context.Context, worker *ResolvedWorker) {
	p.waitGroup.Add(1)

	workerExecution := p.workerExecution(worker)

	go func(ctx context.Context, workerExecution *WorkerExecution) {
		defer func() {
			p.waitGroup.Done()

			if r := recover(); r != nil {
				workerExecution.
					SetStatus(Stopped).
					AddEvent(fmt.Sprintf(
						"stopping execution %d/%d with recovered panic: %s",
						workerExecution.CurrentExecutionAttempt(),
						workerExecution.MaxExecutionsAttempts(),
						r,
					))

				if workerExecution.CurrentExecutionAttempt() < workerExecution.MaxExecutionsAttempts() {
					workerExecution.AddEvent("restarting after panic recovery")

					p.startResolvedWorker(ctx, worker)
				} else {
					workerExecution.AddEvent("max execution attempts reached")
				}
			}
		}()

		if workerExecution.CurrentExecutionAttempt() == 0 && workerExecution.DeferredStartThreshold() > 0 {
			workerExecution.AddEvent(fmt.Sprintf(
				"deferring execution for %d seconds",
				workerExecution.DeferredStartThreshold(),
			))

			time.Sleep(time.Duration(workerExecution.DeferredStartThreshold()) * time.Second)
		}

		workerExecution.
			SetStatus(Running).
			SetCurrentExecutionAttempt(workerExecution.CurrentExecutionAttempt() + 1).
			AddEvent(fmt.Sprintf(
				"starting execution %d/%d",
				workerExecution.CurrentExecutionAttempt(),
				workerExecution.MaxExecutionsAttempts(),
			))

		if err := worker.Implementation().Run(ctx); err != nil {
			workerExecution.
				SetStatus(Stopped).
				AddEvent(fmt.Sprintf(
					"stopping execution %d/%d with error: %v",
					workerExecution.CurrentExecutionAttempt(),
					workerExecution.MaxExecutionsAttempts(),
					err.Error(),
				))

			if workerExecution.CurrentExecutionAttempt() < workerExecution.MaxExecutionsAttempts() {
				workerExecution.AddEvent("restarting after error")

				p.startResolvedWorker(ctx, worker)
			} else {
				workerExecution.AddEvent("max execution attempts reached")
			}
		} else {
			workerExecution.
				SetStatus(Stopped).
				AddEvent(
					fmt.Sprintf(
						"stopping execution %d/%d with success",
						workerExecution.CurrentExecutionAttempt(),
						workerExecution.MaxExecutionsAttempts(),
					))
		}
	}(ctx, workerExecution)
}

func (p *WorkerPool) workerExecution(worker *ResolvedWorker) *WorkerExecution {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	if _, ok := p.executions[worker.Implementation().Name()]; !ok {
		workerOptions := p.options
		for _, opt := range worker.Options() {
			opt(&workerOptions)
		}

		p.executions[worker.Implementation().Name()] = NewWorkerExecution(worker.Implementation().Name(), workerOptions)
	}

	return p.executions[worker.Implementation().Name()]
}

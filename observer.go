package main

import (
	"sync"
)

// WorkerExecutionObserver executions observer
type WorkerExecutionObserver struct {
	mutex      sync.Mutex
	executions map[string]*WorkerExecution
}

func NewWorkerExecutionObserver() *WorkerExecutionObserver {
	return &WorkerExecutionObserver{
		executions: make(map[string]*WorkerExecution),
	}
}

func (o *WorkerExecutionObserver) Observe(worker Worker, status WorkerStatus, message string) *WorkerExecutionObserver {
	o.mutex.Lock()
	defer o.mutex.Unlock()

	if _, ok := o.executions[worker.Name()]; !ok {
		o.executions[worker.Name()] = NewWorkerExecution(worker)
	}

	o.executions[worker.Name()].SetStatus(status)
	o.executions[worker.Name()].AddEvent(message)

	return o
}

func (o *WorkerExecutionObserver) Executions() map[string]*WorkerExecution {
	o.mutex.Lock()
	defer o.mutex.Unlock()

	return o.executions
}

func (o *WorkerExecutionObserver) WorkerExecution(worker Worker) *WorkerExecution {
	o.mutex.Lock()
	defer o.mutex.Unlock()

	if _, ok := o.executions[worker.Name()]; !ok {
		o.executions[worker.Name()] = NewWorkerExecution(worker)
	}

	return o.executions[worker.Name()]
}

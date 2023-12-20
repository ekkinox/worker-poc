package main

type WorkerPoolFactory interface {
	Create(options ...WorkerOption) (*WorkerPool, error)
}

type DefaultWorkerPoolFactory struct{}

func NewDefaultWorkerPoolFactory() WorkerPoolFactory {
	return &DefaultWorkerPoolFactory{}
}

func (f *DefaultWorkerPoolFactory) Create(options ...WorkerOption) (*WorkerPool, error) {
	return NewWorkerPool(options...), nil
}

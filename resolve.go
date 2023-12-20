package main

type ResolvedWorker struct {
	implementation Worker
	options        []WorkerOption
}

func NewResolvedWorker(implementation Worker, options ...WorkerOption) *ResolvedWorker {
	return &ResolvedWorker{
		implementation: implementation,
		options:        options,
	}
}

func (r *ResolvedWorker) Implementation() Worker {
	return r.implementation
}

func (r *ResolvedWorker) Options() []WorkerOption {
	return r.options
}

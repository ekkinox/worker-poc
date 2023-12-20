package main

type WorkerDefinition interface {
	ReturnType() string
	Options() []WorkerOption
}

type workerDefinition struct {
	returnType string
	options    []WorkerOption
}

func NewWorkerDefinition(returnType string, options ...WorkerOption) WorkerDefinition {
	return &workerDefinition{
		returnType: returnType,
		options:    options,
	}
}

func (w *workerDefinition) ReturnType() string {
	return w.returnType
}

func (w *workerDefinition) Options() []WorkerOption {
	return w.options
}

package main

import (
	"fmt"

	"go.uber.org/fx"
)

type WorkerRegistry struct {
	workers     []Worker
	definitions []WorkerDefinition
}

type FxWorkerRegistryParam struct {
	fx.In
	Workers     []Worker           `group:"workers"`
	Definitions []WorkerDefinition `group:"workers-definitions"`
}

func NewFxWorkerRegistry(p FxWorkerRegistryParam) *WorkerRegistry {
	return &WorkerRegistry{
		workers:     p.Workers,
		definitions: p.Definitions,
	}
}

func (r *WorkerRegistry) ResolveCheckerProbesRegistrations() ([]*ResolvedWorker, error) {
	resolvedWorkers := []*ResolvedWorker{}

	for _, definition := range r.definitions {
		implementation, err := r.lookupRegisteredWorker(definition.ReturnType())
		if err != nil {
			return nil, err
		}

		resolvedWorkers = append(
			resolvedWorkers,
			NewResolvedWorker(implementation, definition.Options()...),
		)
	}

	return resolvedWorkers, nil
}

func (r *WorkerRegistry) lookupRegisteredWorker(returnType string) (Worker, error) {
	for _, implementation := range r.workers {
		if GetType(implementation) == returnType {
			return implementation, nil
		}
	}

	return nil, fmt.Errorf("cannot find worker implementation for type %s", returnType)
}

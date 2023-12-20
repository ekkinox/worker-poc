package main

import (
	"context"

	"go.uber.org/fx"
)

const ModuleName = "fx-worker"

var FxWorkerModule = fx.Module(
	ModuleName,
	fx.Provide(
		NewDefaultWorkerPoolFactory,
		NewFxWorkerRegistry,
		NewFxWorkerPool,
		fx.Annotate(
			NewFxWorkerModuleInfo,
			fx.As(new(interface{})),
			fx.ResultTags(`group:"core-module-infos"`),
		),
	),
)

type FxWorkerPoolParam struct {
	fx.In
	LifeCycle fx.Lifecycle
	Factory   WorkerPoolFactory
	Registry  *WorkerRegistry
}

func NewFxWorkerPool(p FxWorkerPoolParam) (*WorkerPool, error) {
	pool, err := p.Factory.Create(WithMaxExecutionsAttempts(2), WithDeferredStartThreshold(10))
	if err != nil {
		return nil, err
	}

	resolvedWorkers, err := p.Registry.ResolveCheckerProbesRegistrations()
	if err != nil {
		return nil, err
	}

	pool.AddResolvedWorkers(resolvedWorkers...)

	p.LifeCycle.Append(fx.Hook{
		OnStart: func(context.Context) error {
			return pool.Start(context.Background())
		},
		OnStop: func(context.Context) error {
			return pool.Stop()
		},
	})

	return pool, nil
}

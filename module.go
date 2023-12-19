package main

import (
	"context"

	"go.uber.org/fx"
)

const ModuleName = "fx-worker"

var FxWorkerModule = fx.Module(
	ModuleName,
	fx.Provide(
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
	Workers   []Worker `group:"workers"`
}

func NewFxWorkerPool(p FxWorkerPoolParam) *WorkerPool {
	// arbitrary value 3, will come form config
	pool := NewWorkerPool(WithExecutionsLimit(3)).AddWorkers(p.Workers...)

	p.LifeCycle.Append(fx.Hook{
		OnStart: func(context.Context) error {
			return pool.Start(context.Background())
		},
		OnStop: func(context.Context) error {
			return pool.Stop()
		},
	})

	return pool
}

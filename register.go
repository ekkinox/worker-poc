package main

import "go.uber.org/fx"

func AsWorker(w any, options ...WorkerOption) fx.Option {
	return fx.Options(
		fx.Provide(
			fx.Annotate(
				w,
				fx.As(new(Worker)),
				fx.ResultTags(`group:"workers"`),
			),
		),
		fx.Supply(
			fx.Annotate(
				NewWorkerDefinition(GetReturnType(w), options...),
				fx.As(new(WorkerDefinition)),
				fx.ResultTags(`group:"workers-definitions"`),
			),
		),
	)
}

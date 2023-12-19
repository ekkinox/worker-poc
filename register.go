package main

import "go.uber.org/fx"

func AsWorker(worker any) fx.Option {
	return fx.Provide(
		fx.Annotate(
			worker,
			fx.As(new(Worker)),
			fx.ResultTags(`group:"workers"`),
		),
	)
}

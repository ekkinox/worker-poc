package main

import (
	"time"

	"github.com/ekkinox/worker-poc/workers"
	"go.uber.org/fx"
)

func main() {
	var pool *WorkerPool

	app := fx.New(
		// register worker module
		FxWorkerModule,
		// register workers
		AsWorker(workers.NewSuccessWorker),
		AsWorker(workers.NewErrorWorker),
		AsWorker(workers.NewPanicWorker),
		// periodic observation
		fx.Invoke(func(pool *WorkerPool) {
			go func(pool *WorkerPool) {
				for {
					time.Sleep(3 * time.Second)

					printObservation("periodical observation", pool)
				}

			}(pool)
		}),
		// pool extraction
		fx.Populate(&pool),
		// to let time to workers to finish
		fx.StopTimeout(3*time.Second),
	)

	app.Run()

	time.Sleep(1 * time.Second)

	printObservation("last observation", pool)
}

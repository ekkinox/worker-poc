package main

import (
	"fmt"
	"time"

	"github.com/ekkinox/worker-poc/workers"
	"go.uber.org/fx"
)

func main() {
	var pool *WorkerPool

	app := fx.New(
		fx.StopTimeout(3*time.Second),
		FxWorkerModule,
		AsWorker(workers.NewSuccessWorker),
		AsWorker(workers.NewErrorWorker),
		AsWorker(workers.NewPanicWorker),
		fx.Populate(&pool),
		fx.Invoke(func(pool *WorkerPool) {
			go func() {
				for {
					time.Sleep(2 * time.Second)

					printObservation("periodical observation", pool)
				}

			}()
		}),
	)

	app.Run()

	time.Sleep(2 * time.Second)

	printObservation("last observation", pool)

}

func printObservation(title string, pool *WorkerPool) {
	fmt.Printf("****\n%s:\n", title)

	for name, execution := range pool.Observer().Executions() {
		fmt.Println("")
		fmt.Printf("%s: %s (%d events)\n", name, execution.Status(), len(execution.Events()))
		for _, event := range execution.Events() {
			fmt.Printf("- message: %s, time: %s\n", event.Message(), event.Timestamp())
		}
	}

	fmt.Println("****\n\n")
}

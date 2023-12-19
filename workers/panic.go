package workers

import (
	"context"
	"fmt"
	"time"
)

type PanicWorker struct{}

func NewPanicWorker() *PanicWorker {
	return &PanicWorker{}
}

func (w *PanicWorker) Name() string {
	return "PanicWorker"
}

func (w *PanicWorker) Run(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("\n####### => ", w.Name(), "is stopping\n")
			time.Sleep(1 * time.Second) // simulate work
			fmt.Println("\n####### => ", w.Name(), "is stopped\n")
			return nil
		default:
			fmt.Println("\n####### => ", w.Name(), "is running\n")
			time.Sleep(1 * time.Second) // simulate work
			panic("custom panic")
		}
	}
}

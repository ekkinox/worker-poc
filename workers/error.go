package workers

import (
	"context"
	"fmt"
	"time"
)

type ErrorWorker struct{}

func NewErrorWorker() *ErrorWorker {
	return &ErrorWorker{}
}

func (w *ErrorWorker) Name() string {
	return "ErrorWorker"
}

func (w *ErrorWorker) Run(ctx context.Context) error {
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
			return fmt.Errorf("custom error from %s worker", w.Name())
		}
	}
}

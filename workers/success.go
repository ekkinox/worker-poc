package workers

import (
	"context"
	"fmt"
	"time"
)

type SuccessWorker struct{}

func NewSuccessWorker() *SuccessWorker {
	return &SuccessWorker{}
}

func (w *SuccessWorker) Name() string {
	return "SuccessWorker"
}

func (w *SuccessWorker) Run(ctx context.Context) error {
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
		}
	}
}

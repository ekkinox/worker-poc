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
	return "success"
}

func (w *SuccessWorker) Run(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			time.Sleep(1 * time.Second) // simulate work
			fmt.Println(w.Name(), "stopped")
			return nil
		default:
			time.Sleep(1 * time.Second) // simulate work
			fmt.Println("\n####### => ", w.Name(), "is running\n")
			time.Sleep(1 * time.Second) // simulate work
		}
	}
}

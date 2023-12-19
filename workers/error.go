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
	return "error"
}

func (w *ErrorWorker) Run(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			time.Sleep(1 * time.Second) // simulate work
			return nil
		default:
			time.Sleep(1 * time.Second) // simulate work
			return fmt.Errorf("custom error from %s worker", w.Name())
		}
	}
}
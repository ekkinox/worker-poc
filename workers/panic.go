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
	return "panic"
}

func (w *PanicWorker) Run(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			time.Sleep(1 * time.Second) // simulate work
			return nil
		default:
			time.Sleep(1 * time.Second) // simulate work
			panic(fmt.Sprintf("panic form %s worker", w.Name()))
		}
	}
}

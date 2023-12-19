package main

import (
	"context"
)

// Worker interface
type Worker interface {
	Name() string
	Run(ctx context.Context) error
}

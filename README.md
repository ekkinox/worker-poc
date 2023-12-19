# Worker module PoC

> PoC for a [Fx](https://uber-go.github.io/fx/) compatible worker module.

<!-- TOC -->
* [Usage](#usage)
* [Worker registration](#worker-registration)
<!-- TOC -->

## Usage

Simply run:
```go
go run .
```

After a while, you can `ctrl+c` to see graceful shutdown.

## Worker registration

To register a worker, you need to implement the [Worker](worker.go) interface.

```go
package workers

import (
	"context"
	"time"
)

type MyWorker struct{}

func NewMyWorker() *MyWorker {
	return &MyWorker{}
}

func (w *MyWorker) Name() string {
	return "my worker"
}

func (w *MyWorker) Run(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			return nil
		default:
			time.Sleep(1 * time.Second) // simulate work
		}
	}
}
```

You can then `register` it in Fx using the `AsWorker()` function in [main.go](main.go):

```go
fx.New(
	// ...
	AsWorker(workers.NewMyWorker),
	// ...
)
```

Your worker dependencies will be autowired.
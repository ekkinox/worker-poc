package main

type Options struct {
	DeferredStartThreshold int
	MaxExecutionsAttempts  int
}

type WorkerOption func(o *Options)

func DefaultWorkerOptions() Options {
	return Options{
		DeferredStartThreshold: 0,
		MaxExecutionsAttempts:  1,
	}
}

func WithDeferredStartThreshold(t int) WorkerOption {
	return func(o *Options) {
		o.DeferredStartThreshold = t
	}
}

func WithMaxExecutionsAttempts(l int) WorkerOption {
	return func(o *Options) {
		o.MaxExecutionsAttempts = l
	}
}

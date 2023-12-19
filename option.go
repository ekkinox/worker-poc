package main

type Options struct {
	ExecutionsLimit int
}

type WorkerPoolOption func(o *Options)

func DefaultWorkerPoolOptions() Options {
	return Options{
		ExecutionsLimit: 1,
	}
}

func WithExecutionsLimit(l int) WorkerPoolOption {
	return func(o *Options) {
		o.ExecutionsLimit = l
	}
}

package main

type FxWorkerModuleInfo struct {
	pool *WorkerPool
}

func NewFxWorkerModuleInfo(pool *WorkerPool) *FxWorkerModuleInfo {
	return &FxWorkerModuleInfo{
		pool: pool,
	}
}

func (i *FxWorkerModuleInfo) Name() string {
	return ModuleName
}

func (i *FxWorkerModuleInfo) Data() map[string]interface{} {
	return map[string]interface{}{
		"workers": i.pool.Report(),
	}
}

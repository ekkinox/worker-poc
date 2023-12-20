package main

import (
	"sync"
	"time"
)

// WorkerExecutionEvent execution
type WorkerExecutionEvent struct {
	message   string
	timestamp time.Time
}

func NewWorkerExecutionEvent(message string) *WorkerExecutionEvent {
	return &WorkerExecutionEvent{
		message:   message,
		timestamp: time.Now(),
	}
}

func (e *WorkerExecutionEvent) Message() string {
	return e.message
}

func (e *WorkerExecutionEvent) Timestamp() time.Time {
	return e.timestamp
}

// WorkerExecution execution
type WorkerExecution struct {
	mutex                   sync.Mutex
	name                    string
	status                  WorkerStatus
	currentExecutionAttempt int
	maxExecutionsAttempts   int
	deferredStartThreshold  int
	events                  []*WorkerExecutionEvent
}

func NewWorkerExecution(name string, options Options) *WorkerExecution {
	return &WorkerExecution{
		name:                    name,
		status:                  Unknown,
		currentExecutionAttempt: 0,
		maxExecutionsAttempts:   options.MaxExecutionsAttempts,
		deferredStartThreshold:  options.DeferredStartThreshold,
		events:                  []*WorkerExecutionEvent{},
	}
}

func (e *WorkerExecution) Name() string {
	return e.name
}

func (e *WorkerExecution) Status() WorkerStatus {
	e.mutex.Lock()
	defer e.mutex.Unlock()

	return e.status
}

func (e *WorkerExecution) SetStatus(status WorkerStatus) *WorkerExecution {
	e.mutex.Lock()
	defer e.mutex.Unlock()

	e.status = status

	return e
}

func (e *WorkerExecution) CurrentExecutionAttempt() int {
	e.mutex.Lock()
	defer e.mutex.Unlock()

	return e.currentExecutionAttempt
}

func (e *WorkerExecution) SetCurrentExecutionAttempt(current int) *WorkerExecution {
	e.mutex.Lock()
	defer e.mutex.Unlock()

	e.currentExecutionAttempt = current

	return e
}

func (e *WorkerExecution) MaxExecutionsAttempts() int {
	e.mutex.Lock()
	defer e.mutex.Unlock()

	return e.maxExecutionsAttempts
}

func (e *WorkerExecution) SetMaxExecutionsAttempts(max int) *WorkerExecution {
	e.mutex.Lock()
	defer e.mutex.Unlock()

	e.maxExecutionsAttempts = max

	return e
}

func (e *WorkerExecution) DeferredStartThreshold() int {
	e.mutex.Lock()
	defer e.mutex.Unlock()

	return e.deferredStartThreshold
}

func (e *WorkerExecution) SetDeferredStartThreshold(threshold int) *WorkerExecution {
	e.mutex.Lock()
	defer e.mutex.Unlock()

	e.deferredStartThreshold = threshold

	return e
}

func (e *WorkerExecution) Events() []*WorkerExecutionEvent {
	e.mutex.Lock()
	defer e.mutex.Unlock()

	return e.events
}

func (e *WorkerExecution) AddEvent(message string) *WorkerExecution {
	e.mutex.Lock()
	defer e.mutex.Unlock()

	e.events = append(e.events, NewWorkerExecutionEvent(message))

	return e
}

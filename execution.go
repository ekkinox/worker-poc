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
	mutex   sync.Mutex
	name    string
	status  WorkerStatus
	current int
	limit   int
	events  []*WorkerExecutionEvent
}

func NewWorkerExecution(worker Worker, limit int) *WorkerExecution {
	return &WorkerExecution{
		name:    worker.Name(),
		status:  Unknown,
		current: 0,
		limit:   limit,
		events:  []*WorkerExecutionEvent{},
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

func (e *WorkerExecution) Current() int {
	e.mutex.Lock()
	defer e.mutex.Unlock()

	return e.current
}

func (e *WorkerExecution) SetCurrent(current int) *WorkerExecution {
	e.mutex.Lock()
	defer e.mutex.Unlock()

	e.current = current

	return e
}

func (e *WorkerExecution) Limit() int {
	e.mutex.Lock()
	defer e.mutex.Unlock()

	return e.limit
}

func (e *WorkerExecution) SetLimit(limit int) *WorkerExecution {
	e.mutex.Lock()
	defer e.mutex.Unlock()

	e.limit = limit

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

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
	mutex  sync.Mutex
	name   string
	status WorkerStatus
	events []*WorkerExecutionEvent
}

func NewWorkerExecution(worker Worker) *WorkerExecution {
	return &WorkerExecution{
		name:   worker.Name(),
		status: Unknown,
		events: []*WorkerExecutionEvent{},
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

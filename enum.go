package main

type WorkerStatus int

const (
	Unknown WorkerStatus = iota
	Running
	Stopped
)

func (s WorkerStatus) String() string {
	switch s {
	case Running:
		return "running"
	case Stopped:
		return "stopped"
	default:
		return "unknown"
	}
}

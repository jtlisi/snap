package core

import "time"

type TaskState int

const (
	TaskStopped TaskState = iota
	TaskSpinning
	TaskFiring
)

type Task interface {
	Id() uint64
	// Status() WorkflowState TODO, switch to string
	State() TaskState
	HitCount() uint
	MissedCount() uint
	LastRunTime() time.Time
	CreationTime() time.Time
	DeadlineDuration() time.Duration
	SetDeadlineDuration(time.Duration)
	Option(...TaskOption) TaskOption
}

type TaskOption func(Task) TaskOption

// TaskDeadlineDuration sets the tasks deadline.
// The deadline is the amount of time that can pass before a worker begins
// processing the tasks collect job.
func TaskDeadlineDuration(v time.Duration) TaskOption {
	return func(t Task) TaskOption {
		previous := t.DeadlineDuration()
		t.SetDeadlineDuration(v)
		return TaskDeadlineDuration(previous)
	}
}

type TaskErrors interface {
	Errors() []error
}
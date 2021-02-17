package kuntul

type Adapter interface {
	Lock(task *Task) error
	Unlock() error
}

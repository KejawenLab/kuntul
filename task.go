package kuntul

import "time"

type Task struct {
	ID         string
	Cmd        func(job *Job)
	Estimation time.Duration
	Schedule   string
}

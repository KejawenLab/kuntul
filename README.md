# Kuntul

Kuntul is distributed task scheduler

## Usage

```go
package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/KejawenLab/kuntul"
	"github.com/KejawenLab/kuntul/adapters"
)

func main() {
	task := &kuntul.Task{
		ID: "abc",
		Cmd: func(job *kuntul.Job) {
			fmt.Println(time.Now().Clock())

			job.Done()
		},
		Schedule:   "* * * * * *",
		Estimation: 3 * time.Second,
	}

	job := kuntul.NewJob(adapters.NewRedisAdapter("localhost:6379"))
	job.Add(task)

	go job.Start()

	sig := make(chan os.Signal)
	signal.Notify(sig, os.Interrupt, os.Kill)
	<-sig
}

```

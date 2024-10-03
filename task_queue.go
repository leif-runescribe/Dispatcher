package main

import (
	"fmt"
	"sync"
)

type TaskQueue struct {
	jobs  chan Job
	mu    sync.Mutex
	jobID int
}

func newTaskQueue(bufferSize int) *TaskQueue {
	return &TaskQueue{
		jobs:  make(chan Job, bufferSize),
		jobID: 0,
	}

}

func (x *TaskQueue) AddJob(name string, desc string) Job {
	x.mu.Lock()
	x.jobID++

	var j Job = Job{
		ID:   x.jobID,
		name: name,
		desc: desc,
	}

	x.jobs <- j
	x.mu.Unlock()

	fmt.Printf("A job was just added to the taskqueue with the JobID: %v", x.jobID)
	return j
}

func (q *TaskQueue) getJob() Job {
	job := <-q.jobs
	fmt.Printf("Fetched job: %v", job)
	return job
}

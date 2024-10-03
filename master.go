package main

import (
	"fmt"
	"sync"
)


type MasterServer struct{
	tq *TaskQueue
	workers []chan Job
	mu sync.Mutex
}

func newMasterServer(qSize int) *MasterServer{
	return &MasterServer{
		tq: newTaskQueue(qSize),
		workers: make([]chan Job, 0),
	}
}

func (lilmaster *MasterServer) registerWorker(worker chan Job){
	lilmaster.mu.Lock()
	lilmaster.workers = append(lilmaster.workers, worker)
	lilmaster.mu.Unlock()s
	fmt.Println("Just added a worker!")
}

func (m *MasterServer) DispatchJobs(){
	go func(){
		for{
			job := m.tq.getJob()
			for _, worker := range m.workers{
			worker <- job
			break
			}
		}	
		
	}()
}

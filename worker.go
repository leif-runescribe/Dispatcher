package main

type Worker struct {
	ID    int
	JobCh chan Job
}

func newWorker(id int) *Worker {
	return &Worker{
		ID:    id,
		JobCh: make(chan Job),
	}
}

func (w *Worker) startWorker(master *MasterServer) {
	master.regis
}

func main() {

}

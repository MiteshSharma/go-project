package eventdispatcher

import (
	"github.com/MiteshSharma/project/bus"
	"github.com/MiteshSharma/project/logger"
)

type Worker struct {
	Id           string
	EventChannel chan Event
	WorkerQueue  chan Worker
	Quit         chan bool
	Log          logger.Logger
	Bus          bus.Bus
}

func NewWorker(id string, taskWorkerQueue chan Worker, log logger.Logger, bus bus.Bus) *Worker {
	worker := &Worker{
		Id:           id,
		EventChannel: make(chan Event),
		WorkerQueue:  taskWorkerQueue,
		Quit:         make(chan bool),
		Log:          log,
		Bus:          bus,
	}
	return worker
}

func (w *Worker) Start() {
	go func() {
		for {
			// Adding worker in worker queue
			w.WorkerQueue <- *w
			select {
			case event := <-w.EventChannel:
				// Dispatch work from here
				w.Bus.Publish(event.Name, event.Message)
			case <-w.Quit:
				// Stop this worker
				return
			}
		}
	}()
}

func (w *Worker) Stop() {
	go func() {
		w.Quit <- true
	}()
}

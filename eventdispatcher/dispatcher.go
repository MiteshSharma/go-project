package eventdispatcher

import (
	"github.com/MiteshSharma/project/bus"
	"github.com/MiteshSharma/project/logger"
	uuid "github.com/satori/go.uuid"
)

type Dispatcher struct {
	NumWorker        int
	EventQueue       chan Event
	EventWorkerQueue chan Worker
	Workers          []*Worker
	Quit             chan bool
	Log              logger.Logger
	Bus              bus.Bus
}

func NewDispatcher(numWorker int, EventQueue chan Event, logger logger.Logger, bus bus.Bus) *Dispatcher {
	dispatcher := &Dispatcher{
		NumWorker:  numWorker,
		EventQueue: EventQueue,
		Workers:    make([]*Worker, numWorker),
		Quit:       make(chan bool),
		Log:        logger,
		Bus:        bus,
	}
	return dispatcher
}

func (d *Dispatcher) Start() {
	d.EventWorkerQueue = make(chan Worker, d.NumWorker)

	for count := 0; count < d.NumWorker; count++ {
		uid := uuid.NewV4()
		worker := NewWorker(uid.String(), d.EventWorkerQueue, d.Log, d.Bus)
		d.Workers[count] = worker
		d.Log.Info("Worked created for event handling with id :" + string(worker.Id))
		worker.Start()
	}

	go func() {
		var event Event
		for {
			select {
			case event = <-d.EventQueue:
				go func() {
					var worker = <-d.EventWorkerQueue
					d.Log.Info("Event passed to worker with id :" + string(worker.Id))
					worker.EventChannel <- event
				}()
			case <-d.Quit:
				for count := 0; count < d.NumWorker; count++ {
					d.Workers[count].Quit <- true
				}
				return
			}
		}
	}()
}

func (d *Dispatcher) Stop() {
	go func() {
		d.Quit <- true
	}()
}

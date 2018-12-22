package eventdispatcher

import (
	"errors"

	"github.com/MiteshSharma/project/bus"

	"github.com/MiteshSharma/project/logger"
)

type EventDispatcher struct {
	EventQueue  chan Event
	Dispatcher  *Dispatcher
	Log         logger.Logger
	Bus         bus.Bus
	queueSize   int
	workerCount int
}

func NewEventDispatcher(log logger.Logger, bus bus.Bus, queueSize int, workerCount int) *EventDispatcher {
	eventDispatcher := &EventDispatcher{
		Log:         log,
		Bus:         bus,
		queueSize:   queueSize,
		workerCount: workerCount,
	}
	eventDispatcher.Start()
	return eventDispatcher
}

func (ed *EventDispatcher) Start() {
	ed.EventQueue = make(chan Event, ed.queueSize)
	ed.Dispatcher = NewDispatcher(ed.workerCount, ed.EventQueue, ed.Log, ed.Bus)
	ed.Dispatcher.Start()
}

func (ed *EventDispatcher) Send(event Event) error {
	if event.Name == "" {
		return errors.New("event must have a name which define who is listener")
	}
	go func() {
		ed.EventQueue <- event
	}()
	return nil
}

func (ed *EventDispatcher) Stop() {
	ed.Dispatcher.Stop()
}

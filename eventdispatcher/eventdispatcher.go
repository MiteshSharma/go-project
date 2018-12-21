package eventdispatcher

import (
	"errors"

	"github.com/MiteshSharma/project/logger"
)

type EventDispatcher struct {
	EventQueue  chan Event
	Dispatcher  *Dispatcher
	Log         logger.Logger
	queueSize   int
	workerCount int
}

func NewEventDispatcher(log logger.Logger, queueSize int, workerCount int) *EventDispatcher {
	eventDispatcher := &EventDispatcher{
		Log:         log,
		queueSize:   queueSize,
		workerCount: workerCount,
	}
	eventDispatcher.Start()
	return eventDispatcher
}

func (ed *EventDispatcher) Start() {
	ed.EventQueue = make(chan Event, ed.queueSize)
	ed.Dispatcher = NewDispatcher(ed.workerCount, ed.EventQueue, ed.Log)
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

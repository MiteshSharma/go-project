package bi

import (
	"errors"

	"github.com/MiteshSharma/project/eventdispatcher"
)

// BiEvent struct to send all BI events
type BiEventHandler struct {
	EventDispatcher *eventdispatcher.EventDispatcher
}

func NewBiEventHandler(eventdispatcher *eventdispatcher.EventDispatcher) *BiEventHandler {
	biEventHandler := &BiEventHandler{
		EventDispatcher: eventdispatcher,
	}
	return biEventHandler
}

func (bi BiEventHandler) Send(eventName string, data map[string]interface{}) error {
	if eventName == "" {
		return errors.New("event name must be provided")
	}
	if data == nil {
		data = make(map[string]interface{})
	}
	data["eventName"] = eventName
	event := eventdispatcher.Event{
		Name:    "bi",
		Message: data,
	}
	bi.EventDispatcher.Send(event)
	return nil
}

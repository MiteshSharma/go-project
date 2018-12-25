package bi

// EventHandler struct to send all BI events
type EventHandler interface {
	Send(eventName string, data map[string]interface{}) error
}

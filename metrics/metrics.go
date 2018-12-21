package metrics

type Metrics interface {
	RequestReceived(code string, method string)
}

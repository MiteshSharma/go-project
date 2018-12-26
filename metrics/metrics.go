package metrics

type Metrics interface {
	RequestReceivedDetail(path string, method string, code int, elapsedDuration float64)
}

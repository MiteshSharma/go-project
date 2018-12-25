package metrics

type TestMetrics struct {
}

func NewTestMetrics() Metrics {
	metrics := &TestMetrics{}
	return metrics
}

func (p *TestMetrics) RequestReceived(code string, method string) {
}

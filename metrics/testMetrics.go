package metrics

import "github.com/gorilla/mux"

type TestMetrics struct {
}

func NewTestMetrics() Metrics {
	metrics := &TestMetrics{}
	return metrics
}

func (p *TestMetrics) SetupHttpHandler(router *mux.Router) {

}

func (p *TestMetrics) RequestReceivedDetail(path string, method string, code int, elapsedDuration float64) {
}

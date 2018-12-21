package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

type Prometheus struct {
	RequestCounter *prometheus.CounterVec
}

func NewMetrics() Metrics {
	requestCounter := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "HTTP requests count, partitioned by status code and HTTP method.",
		},
		[]string{"code", "method"},
	)
	metrics := &Prometheus{
		RequestCounter: requestCounter,
	}
	return metrics
}

func (p *Prometheus) RequestReceived(code string, method string) {
	p.RequestCounter.WithLabelValues(code, method).Inc()
}

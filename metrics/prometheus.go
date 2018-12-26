package metrics

import (
	"fmt"

	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
)

type Prometheus struct {
	RequestCounter prometheus.Counter
	RequestSummary *prometheus.SummaryVec
}

func NewMetrics() Metrics {
	requestCounter := prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "HTTP requests total count",
		})
	requestSummary := prometheus.NewSummaryVec(
		prometheus.SummaryOpts{
			Name: "http_request_duration",
			Help: "HTTP request summary with path, status code, method and duration",
		},
		[]string{"path", "statuscode", "method"},
	)
	metrics := &Prometheus{
		RequestCounter: requestCounter,
		RequestSummary: requestSummary,
	}

	prometheus.MustRegister(requestCounter)
	prometheus.MustRegister(requestSummary)
	return metrics
}

func (p *Prometheus) SetupHttpHandler(router *mux.Router) {
	router.Handle("", promhttp.Handler()).Methods("GET")
}

func (p *Prometheus) RequestReceivedDetail(path string, method string, code int, elapsedDuration float64) {
	p.RequestCounter.Inc()
	p.RequestSummary.WithLabelValues(path, fmt.Sprintf("%d", code), method).Observe(elapsedDuration)
}

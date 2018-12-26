package metrics

import (
	"github.com/gorilla/mux"
)

type Metrics interface {
	SetupHttpHandler(router *mux.Router)
	RequestReceivedDetail(path string, method string, code int, elapsedDuration float64)
}

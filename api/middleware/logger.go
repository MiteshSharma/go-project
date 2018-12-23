package middleware

import (
	"net/http"
	"time"

	"github.com/MiteshSharma/project/model"

	"github.com/MiteshSharma/project/logger"

	"github.com/felixge/httpsnoop"
)

// LoggerMiddleware struct
type LoggerMiddleware struct {
	Logger logger.Logger
}

// NewLoggerMiddleware function returns instance of logger middleware
func NewLoggerMiddleware(logger logger.Logger) *LoggerMiddleware {
	loggerMiddleware := &LoggerMiddleware{
		Logger: logger,
	}
	loggerMiddleware.Init()
	return loggerMiddleware
}

// Init function to init anything required for middleware
func (lm *LoggerMiddleware) Init() {
}

// GetMiddlewareHandler function returns middleware used to log requests
func (lm *LoggerMiddleware) GetMiddlewareHandler() func(http.ResponseWriter, *http.Request, http.HandlerFunc) {
	return func(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

		metrix := httpsnoop.CaptureMetrics(next, rw, r)
		requestID := rw.Header().Get(model.HEADER_REQUEST_ID)
		lm.Logger.Info("Request handling completed from logger middleware ", logger.String("RequestID", requestID),
			logger.String("Host", r.Host), logger.String("Method", r.Method), logger.String("Request", r.RequestURI),
			logger.String("RemoteAddress", r.RemoteAddr), logger.String("Referer", r.Referer()),
			logger.String("UserAgent", r.UserAgent()), logger.Int("StatusCode", metrix.Code),
			logger.Int("Duration", int(metrix.Duration/time.Millisecond)))
	}
}

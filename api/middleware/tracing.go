package middleware

import (
	"net/http"

	"github.com/MiteshSharma/project/logger"
	"github.com/MiteshSharma/project/model"
	"github.com/gorilla/context"
	opentracing "github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	zipkin "github.com/openzipkin/zipkin-go-opentracing"
)

const (
	debug         = false
	sameSpan      = false
	traceID128Bit = true // Tracer generate 128 bit traceID
)

// ZipkinMiddleware struct
type ZipkinMiddleware struct {
	ServiceHostPort string
	ServiceName     string
	Logger          logger.Logger
	tracer          opentracing.Tracer
}

// NewZipkinMiddleware function returns instance of zipkin middleware
func NewZipkinMiddleware(logger logger.Logger, serviceName string, zipkinConfig model.ZipkinConfig) *ZipkinMiddleware {
	zipkinMiddleware := &ZipkinMiddleware{
		ServiceHostPort: (zipkinConfig.Host + ":" + zipkinConfig.Port),
		ServiceName:     serviceName,
		Logger:          logger,
	}
	zipkinMiddleware.Init()
	return zipkinMiddleware
}

// Init function to init request details for zipkin middleware
func (zm *ZipkinMiddleware) Init() {
	zipkinHTTPEndpoint := zm.ServiceHostPort + "/api/v1/spans"
	zm.Logger.Info("Init zipkin middleware ", logger.String("zipkinEndpoint", zipkinHTTPEndpoint))
	collector, err := zipkin.NewHTTPCollector(zipkinHTTPEndpoint)
	if err != nil {
		zm.Logger.Error("Creating collector failed in zipkin middleware ", logger.Error(err))
	}
	recorder := zipkin.NewRecorder(collector, debug, zm.ServiceHostPort, zm.ServiceName)
	tracer, err := zipkin.NewTracer(
		recorder,
		zipkin.ClientServerSameSpan(sameSpan),
		zipkin.TraceID128Bit(traceID128Bit),
	)
	if err != nil {
		zm.Logger.Error("Creating tracer failed ", logger.Error(err))
	}
	opentracing.InitGlobalTracer(tracer)
	zm.tracer = tracer
}

// GetMiddlewareHandler function returns middleware used to trace requests
func (zm *ZipkinMiddleware) GetMiddlewareHandler() func(http.ResponseWriter, *http.Request, http.HandlerFunc) {
	return func(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		wireContext, err := zm.tracer.Extract(
			opentracing.TextMap,
			opentracing.HTTPHeadersCarrier(r.Header),
		)
		uuid := context.Get(r, "uuid")
		if err != nil {
			zm.Logger.Debug("Error encountered while trying to extract span ", logger.String("uuid", uuid.(string)), logger.Error(err))
			next(rw, r)
		} else {
			span := zm.tracer.StartSpan(r.URL.Path, ext.RPCServerOption(wireContext))
			defer span.Finish()
			ctx := opentracing.ContextWithSpan(r.Context(), span)
			r = r.WithContext(ctx)
			next(rw, r)
		}
	}
}

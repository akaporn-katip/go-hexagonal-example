package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/propagation"
	semconv "go.opentelemetry.io/otel/semconv/v1.37.0"
	"go.opentelemetry.io/otel/trace"
)

type TracingMiddleware struct {
	tracer trace.Tracer
}

func NewTraceMiddleware(serviceName string) *TracingMiddleware {
	return &TracingMiddleware{
		tracer: otel.Tracer(serviceName),
	}
}

func (t *TracingMiddleware) Handle() gin.HandlerFunc {
	return func(c *gin.Context) {
		r := c.Request
		ctx := otel.GetTextMapPropagator().Extract(
			c,
			propagation.HeaderCarrier(r.Header),
		)

		route := c.FullPath()
		if route == "" {
			route = "not_found"
		}

		// target := r.URL.Path
		spanName := r.Method + " " + route
		ctx, span := t.tracer.Start(ctx, spanName, trace.WithSpanKind(trace.SpanKindServer),
			trace.WithAttributes(
				semconv.HTTPRequestMethodKey.String(r.Method),
				semconv.HTTPRoute(route),            // Pattern: /api/customer/:id
				semconv.URLPath(r.URL.Path),         // Specific: /api/customer/123
				semconv.ClientAddress(c.ClientIP()), // Cleaner IP extraction
				semconv.UserAgentOriginal(r.UserAgent()),
				semconv.HTTPRequestBodySize(int(r.ContentLength)),
			),
		)
		defer span.End()

		c.Header("X-Trace-Id", span.SpanContext().TraceID().String())
		c.Header("X-Span-Id", span.SpanContext().SpanID().String())
		c.Request = c.Request.WithContext(ctx)
		c.Next()

		rw := c.Writer
		span.SetAttributes(
			semconv.HTTPResponseStatusCode(rw.Status()),
			semconv.HTTPResponseBodySize(rw.Size()),
		)

		if rw.Status() >= 400 {
			span.SetStatus(codes.Error, http.StatusText(rw.Status()))
		} else {
			span.SetStatus(codes.Ok, "")
		}

	}
}

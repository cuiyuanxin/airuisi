package middleware

import (
	"context"

	"github.com/cuiyuanxin/airuisi/global"
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
)

func Tarcing() func(c *gin.Context) {
	return func(c *gin.Context) {
		var ctx context.Context
		span := opentracing.SpanFromContext(c.Request.Context())
		if span != nil {
			span, ctx = opentracing.StartSpanFromContextWithTracer(c.Request.Context(), global.Tracet, c.Request.URL.Path, opentracing.ChildOf(span.Context()))
		} else {
			span, ctx = opentracing.StartSpanFromContextWithTracer(c.Request.Context(), global.Tracet, c.Request.URL.Path)
		}

		defer span.Finish()

		var tracelD string
		var SpanlD string
		var spanContext = span.Context()
		switch spanContext.(type) {
		case jaeger.SpanContext:
			tracelD = spanContext.(jaeger.SpanContext).TraceID().String()
			SpanlD = spanContext.(jaeger.SpanContext).SpanID().String()
		}
		c.Set("X-Trace-ID", tracelD)
		c.Set("X-Span-ID", SpanlD)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

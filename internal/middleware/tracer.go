package middleware

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	"goblog/global"
)

// 链路追踪 中间件
func Tracing() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx context.Context
		span := opentracing.SpanFromContext(c.Request.Context())
		if span != nil {
			span, ctx = opentracing.StartSpanFromContextWithTracer(
				c.Request.Context(), global.Tracer, c.Request.URL.Path, opentracing.ChildOf(span.Context()))
		} else {
			span, ctx = opentracing.StartSpanFromContextWithTracer(c.Request.Context(), global.Tracer, c.Request.URL.Path)
		}
		defer span.Finish()
		// 链路追踪信息注入 context
		var traceID string
		var spanID string
		var spanContext = span.Context()
		jsc := spanContext.(jaeger.SpanContext)
		traceID = jsc.TraceID().String()
		spanID = jsc.SpanID().String()
		// 这里的链路追踪 id 由 opentracing 服务提供
		c.Set("X-Trace-ID", traceID)
		c.Set("X-Span-ID", spanID)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

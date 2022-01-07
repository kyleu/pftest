package httpmetrics

import (
	"github.com/valyala/fasthttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"go.uber.org/zap"
)

var _ propagation.TextMapCarrier = (*headerCarrier)(nil)

type headerCarrier struct {
	h *fasthttp.RequestHeader
}

func (hc headerCarrier) Get(key string) string {
	return string(hc.h.Peek(key))
}

func (hc headerCarrier) Set(key string, value string) {
	hc.h.Set(key, value)
}

func (hc headerCarrier) Keys() []string {
	var keys []string
	hc.h.VisitAll(func(key []byte, _ []byte) {
		keys = append(keys, string(key))
	})
	return keys
}

func ExtractHeaders(rc *fasthttp.RequestCtx, logger *zap.SugaredLogger) *fasthttp.RequestCtx {
	nc := otel.GetTextMapPropagator().Extract(rc, headerCarrier{h: &rc.Request.Header})
	http, ok := nc.(*fasthttp.RequestCtx)
	if !ok {
		logger.Warnf("unable to extract http tracing headers")
		return rc
	}
	return http
}

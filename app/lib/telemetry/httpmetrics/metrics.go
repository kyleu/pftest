package httpmetrics

import (
	"strings"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/valyala/fasthttp"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	"go.opentelemetry.io/otel/trace"
)

type Metrics struct {
	reqCnt  *prometheus.CounterVec
	reqDur  *prometheus.HistogramVec
	reqSize prometheus.Summary
	rspSize prometheus.Summary

	MetricsPath string
}

func NewMetrics(subsystem string) *Metrics {
	ss := strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(subsystem, "-", "_"), "/", "_"), ".", "_")
	m := &Metrics{MetricsPath: defaultMetricPath}
	m.registerHTTPMetrics(ss)
	return m
}

func InjectHTTP(rc *fasthttp.RequestCtx, span trace.Span) {
	span.SetAttributes(
		semconv.HTTPHostKey.String(string(rc.Request.Host())),
		semconv.HTTPMethodKey.String(string(rc.Method())),
		semconv.HTTPURLKey.String(string(rc.Request.RequestURI())),
		semconv.HTTPSchemeKey.String(string(rc.Request.URI().Scheme())),
	)
	if b := rc.Request.Header.Peek("User-Agent"); len(b) > 0 {
		span.SetAttributes(semconv.HTTPUserAgentKey.String(string(b)))
	}
	if b := rc.Request.Header.Peek("Content-Length"); len(b) > 0 {
		span.SetAttributes(semconv.HTTPRequestContentLengthKey.String(string(b)))
	}
	span.SetAttributes(semconv.HTTPAttributesFromHTTPStatusCode(rc.Response.StatusCode())...)
	span.SetStatus(semconv.SpanStatusFromHTTPStatusCode(rc.Response.StatusCode()))
}

func (p *Metrics) registerHTTPMetrics(subsystem string) {
	cOpts := prometheus.CounterOpts{Subsystem: subsystem, Name: "requests_total", Help: "The HTTP request counts processed."}
	p.reqCnt = prometheus.NewCounterVec(cOpts, []string{"code", "method"})

	hOpts := prometheus.HistogramOpts{Subsystem: subsystem, Name: "request_duration_seconds", Help: "The HTTP request duration in seconds."}
	hOpts.Buckets = []float64{.005, .01, .025, .05, .1, .25, .5, 1, 2.5, 5, 10, 15, 20, 30, 40, 50, 60}
	p.reqDur = prometheus.NewHistogramVec(hOpts, []string{"code"})

	reqOpts := prometheus.SummaryOpts{Subsystem: subsystem, Name: "request_size_bytes", Help: "The HTTP request sizes in bytes."}
	p.reqSize = prometheus.NewSummary(reqOpts)

	rspOpts := prometheus.SummaryOpts{Subsystem: subsystem, Name: "response_size_bytes", Help: "The HTTP response sizes in bytes."}
	p.rspSize = prometheus.NewSummary(rspOpts)

	prometheus.MustRegister(p.reqCnt, p.reqDur, p.reqSize, p.rspSize)
}

package observations

import (
	"context"

	rpc "github.com/libp2p/go-libp2p-gorpc"

	"contrib.go.opencensus.io/exporter/jaeger"
	prom "github.com/prometheus/client_golang/prometheus"
)

// PromRegistry is the metrics Registry used by Cluster.
var PromRegistry = prom.NewRegistry()

// SetupMetrics configures and starts stats tooling,
// if enabled.
func SetupMetrics(cfg *MetricsConfig) error { _ = "STUB: not implemented"; return nil }

// JaegerTracer implements ipfscluster.Tracer.
type JaegerTracer struct {
	jaeger *jaeger.Exporter
}

// SetClient no-op.
func (t *JaegerTracer) SetClient(*rpc.Client) {
	_ = "STUB: not implemented"

	// Shutdown the tracer and flush any remaining traces.
	return
}

func (t *JaegerTracer) Shutdown(context.Context) error {
	_ = "STUB: not implemented"
	// nil check for testing, where tracer may not be configured
	return nil
}

// SetupTracing configures and starts tracing tooling,
// if enabled.
func SetupTracing(cfg *TracingConfig) (*JaegerTracer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func setupMetrics(cfg *MetricsConfig) error {
	_ = "STUB: not implemented"
	// setup Prometheus
	return nil
}

// register prometheus with opencensus

// register the metrics views of interest

// setupTracing configures a OpenCensus Tracing exporter for Jaeger.
func setupTracing(cfg *TracingConfig) (*jaeger.Exporter, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// setup Jaeger

// register jaeger with opencensus

// configure tracing

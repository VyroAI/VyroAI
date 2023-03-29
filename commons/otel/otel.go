package otel

import (
	"context"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
	"go.opentelemetry.io/otel/trace"
	"os"
)

var (
	endpoint = "ingest.lightstep.com:443"
)

type Tracer interface {
	NewTracer() trace.Tracer
	InitFiberTrace() *sdktrace.TracerProvider
}

type tracer struct {
	lsToken        string
	lsEnvironment  string
	serviceName    string
	serviceVersion string
}

func InitTracing(serviceName, serviceVersion string) Tracer {
	return &tracer{
		serviceName:    serviceName,
		serviceVersion: serviceVersion,
		lsToken:        os.Getenv("LIGHTSTEP_ACCESSTOKEN"),
		lsEnvironment:  os.Getenv("LIGHTSTEP_ENV"),
	}

}

func (t *tracer) NewExporter(ctx context.Context) (*otlptrace.Exporter, error) {
	var headers = map[string]string{
		"lightstep-access-token": t.lsToken,
	}

	client := otlptracegrpc.NewClient(
		otlptracegrpc.WithEndpoint(endpoint),
		otlptracegrpc.WithHeaders(headers),
	)

	return otlptrace.New(ctx, client)
}

func (t *tracer) NewTraceProvider(exp *otlptrace.Exporter) *sdktrace.TracerProvider {
	resource, rErr :=
		resource.Merge(
			resource.Default(),
			resource.NewWithAttributes(
				semconv.SchemaURL,
				semconv.ServiceNameKey.String(t.serviceName),
				semconv.ServiceVersionKey.String(t.serviceVersion),
				attribute.String("environment", t.lsEnvironment),
			),
		)

	if rErr != nil {
		panic(rErr)
	}

	return sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exp),
		sdktrace.WithResource(resource),
	)

}

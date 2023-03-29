package otel

import (
	"context"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"log"
)

func (t *tracer) InitFiberTrace() *sdktrace.TracerProvider {
	ctx := context.Background()

	exp, err := t.NewExporter(ctx)
	if err != nil {
		log.Fatalf("failed to initialize exporter: %v", err)
	}

	tp := t.NewTraceProvider(exp)

	otel.SetTracerProvider(tp)

	otel.SetTextMapPropagator(
		propagation.NewCompositeTextMapPropagator(
			propagation.TraceContext{},
			propagation.Baggage{},
		),
	)
	//item := tp.Tracer("backend-api", trace.WithInstrumentationVersion("0.1.0"))

	return tp
}

package otel

import (
	"context"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
	"log"
)

func (t *tracer) NewTracer() trace.Tracer {
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

	return tp.Tracer(t.serviceName, trace.WithInstrumentationVersion(t.serviceVersion))

}

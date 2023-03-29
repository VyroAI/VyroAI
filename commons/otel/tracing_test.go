package otel

import (
	"context"
	"fmt"
	"testing"
)

func TestTrace(t *testing.T) {
	ctx := context.Background()

	tracer := InitTracing("testing", "0.1.0")

	trace := tracer.NewTracer()

	ctx, span := trace.Start(ctx, "add")
	defer span.End()

	results := Adding(ctx, 1, 1)

	fmt.Println(results)
}

func Adding(ctx context.Context, num1, num2 int) int {
	return num1 + num2
}

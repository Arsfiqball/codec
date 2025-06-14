package job

import (
	"context"
	"time"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

type ProcessSomething struct {
	ID string `json:"id"`

	// TODO: Add your input fields here
}

func (p ProcessSomething) Kind() string {
	return "module_process_something"
}

type Something struct {
	CompletedAt time.Time

	// TODO: Add your output fields here
}

func (s *service) ProcessSomething(ctx context.Context, input ProcessSomething) (Something, error) {
	attrs := trace.WithAttributes(attribute.String("id", input.ID))

	_, span := s.tracer.Start(ctx, "module/internal/job.Service.ProcessSomething", attrs)
	defer span.End()

	// TODO: Implement your business logic here

	return Something{
		CompletedAt: time.Now(),
	}, nil
}

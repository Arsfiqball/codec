package job

import (
	"context"
	"errors"

	"github.com/Arsfiqball/codec/widget/flame"

	"github.com/gocraft/work"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

type WorkService struct {
	tracer  trace.Tracer
	service Service
}

func NewWorkService(tracer trace.Tracer, service Service) (*WorkService, error) {
	if tracer == nil {
		return nil, errors.New("tracer is required")
	}

	if service == nil {
		return nil, errors.New("service is required")
	}

	return &WorkService{
		tracer:  tracer,
		service: service,
	}, nil
}

func (s *WorkService) ProcessSomething(j *work.Job) error {
	attrs := trace.WithAttributes(attribute.String("job_id", j.ID))

	ctx, span := s.tracer.Start(context.Background(), "module/internal/cross.WorkService.ProcessSomething", attrs)
	defer span.End()

	input := ProcessSomething{
		ID: j.ArgString("id"), // Argument id is not the same as job id
		// TODO: Add your input fields here
	}

	output, err := s.service.ProcessSomething(ctx, input)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "failed to process job")

		return flame.Unexpected(err)
	}

	span.AddEvent("job completed", trace.WithAttributes(
		attribute.String("completed_at", output.CompletedAt.String()),
	))

	return nil
}

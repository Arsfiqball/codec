package scheduler

import (
	"context"
	"errors"

	"go.opentelemetry.io/otel/trace"
)

type FiberService struct {
	tracer  trace.Tracer
	service Service
}

func NewFiberService(tracer trace.Tracer, service Service) (*FiberService, error) {
	if tracer == nil {
		return nil, errors.New("tracer is required")
	}

	if service == nil {
		return nil, errors.New("service is required")
	}

	return &FiberService{
		tracer:  tracer,
		service: service,
	}, nil
}

func (s *FiberService) ScheduleSomething(ctx context.Context, input ScheduleSomethingInput) (ScheduleSomethingOutput, error) {
	ctx, span := s.tracer.Start(ctx, "gadget/internal/scheduler.FiberService.ScheduleSomething") // Don't forget to change the span name
	defer span.End()

	// TODO: Implement your business logic here

	return ScheduleSomethingOutput{}, nil
}

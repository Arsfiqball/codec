package job

import (
	"context"
	"errors"

	"go.opentelemetry.io/otel/trace"
)

type Service interface {
	ProcessSomething(context.Context, ProcessSomething) (Something, error)
}

var _ Service = (*service)(nil)

type service struct {
	tracer trace.Tracer
}

func NewService(
	tracer trace.Tracer,
) (Service, error) {
	if tracer == nil {
		return nil, errors.New("tracer is required")
	}

	return &service{
		tracer: tracer,
	}, nil
}

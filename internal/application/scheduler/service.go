package scheduler

import (
	"context"
	"feature/internal/value/queue"

	"go.opentelemetry.io/otel/trace"
)

type IService interface {
	ResetEverybodyPassword(context.Context) error
}

type Service struct {
	tracer trace.Tracer
	queuer queue.Queuer
}

func NewService(
	tracer trace.Tracer,
	queuer queue.Queuer,
) *Service {
	return &Service{
		tracer: tracer,
		queuer: queuer,
	}
}

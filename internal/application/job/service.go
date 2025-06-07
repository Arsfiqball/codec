package job

import (
	"context"
	"time"

	"github.com/Arsfiqball/codec/internal/value/domain"

	"go.opentelemetry.io/otel/trace"
)

type IService interface {
	ResetEverybodyPassword(ctx context.Context, before time.Time) error
}

type Service struct {
	tracer trace.Tracer
	repo   domain.Repo
}

func NewService(
	tracer trace.Tracer,
	repo domain.Repo,
) *Service {
	return &Service{
		tracer: tracer,
		repo:   repo,
	}
}

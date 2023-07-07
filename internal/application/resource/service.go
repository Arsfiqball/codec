package resource

import (
	"context"
	"feature/internal/value/domain"
	"feature/internal/value/user"

	"go.opentelemetry.io/otel/trace"
)

type IService interface {
	Create(context.Context, domain.Patch, user.Entity) (domain.Entity, error)
	Update(context.Context, domain.Query, domain.Patch, user.Entity) (domain.Entity, error)
	Delete(context.Context, domain.Query, user.Entity) (domain.Entity, error)
	GetOne(context.Context, domain.Query, user.Entity) (domain.Entity, error)
	GetList(context.Context, domain.Query, user.Entity) ([]domain.Entity, error)
	GetStat(context.Context, domain.Query, user.Entity) ([]domain.Stat, error)
}

type Service struct {
	tracer trace.Tracer
	repo   domain.Repo
	event  domain.Event
}

func NewService(
	tracer trace.Tracer,
	repo domain.Repo,
	event domain.Event,
) *Service {
	return &Service{
		tracer: tracer,
		repo:   repo,
		event:  event,
	}
}

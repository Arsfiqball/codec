package resource

import (
	"context"
	"feature/internal/value/domain"
	"feature/internal/value/user"

	"go.opentelemetry.io/otel/trace"
)

type IService interface {
	Create(context.Context, domain.Patch, user.Identity) (domain.Entity, error)
	Update(context.Context, domain.Query, domain.Patch, user.Identity) (domain.Entity, error)
	Delete(context.Context, domain.Query, user.Identity) (domain.Entity, error)
	GetOne(context.Context, domain.Query, user.Identity) (domain.Entity, error)
	GetList(context.Context, domain.Query, user.Identity) ([]domain.Entity, error)
	GetStat(context.Context, domain.Query, user.Identity) ([]domain.Stat, error)
	BulkOps(context.Context, []Ops, user.Identity) ([]OpsResult, error)
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

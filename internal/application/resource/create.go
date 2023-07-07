package resource

import (
	"context"
	"feature/internal/value/domain"
	"feature/internal/value/user"
)

func (s *Service) Create(ctx context.Context, patch domain.Patch, user user.Entity) (domain.Entity, error) {
	ctx, span := s.tracer.Start(ctx, "feature/internal/application/resource.Service.Create")
	defer span.End()

	ent := domain.NewEntity()

	if err := ent.Patch(patch); err != nil {
		return nil, NewError(err, err.Error(), ErrCodeInvalidEntity)
	}

	if err := ent.Validate(); err != nil {
		return nil, NewError(err, err.Error(), ErrCodeInvalidEntity)
	}

	if err := ent.AuthorizeCreate(user); err != nil {
		return nil, NewError(err, err.Error(), ErrCodeUnauthorized)
	}

	ent, err := s.repo.Create(ctx, ent)
	if err != nil {
		return nil, NewError(err, err.Error(), ErrCodeUnknown)
	}

	if err := s.event.Created(ctx, ent); err != nil {
		return nil, NewError(err, err.Error(), ErrCodeUnknown)
	}

	return ent, nil
}

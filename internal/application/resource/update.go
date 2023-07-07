package resource

import (
	"context"
	"feature/internal/value/domain"
	"feature/internal/value/user"
)

func (s *Service) Update(ctx context.Context, query domain.Query, patch domain.Patch, user user.Entity) (domain.Entity, error) {
	ctx, span := s.tracer.Start(ctx, "feature/internal/application/resource.Service.Update")
	defer span.End()

	if err := query.Validate(); err != nil {
		return nil, NewError(err, err.Error(), ErrCodeInvalidQuery)
	}

	old, err := s.repo.GetOne(ctx, query)
	if err != nil {
		return nil, NewError(err, err.Error(), ErrCodeUnknown)
	}

	ent := domain.NewEntityClone(old)

	if err := ent.Patch(patch); err != nil {
		return nil, NewError(err, err.Error(), ErrCodeInvalidEntity)
	}

	if err := ent.Validate(); err != nil {
		return nil, NewError(err, err.Error(), ErrCodeInvalidEntity)
	}

	if err := ent.AuthorizeUpdate(user); err != nil {
		return nil, NewError(err, err.Error(), ErrCodeUnauthorized)
	}

	ent, err = s.repo.Update(ctx, ent)
	if err != nil {
		return nil, NewError(err, err.Error(), ErrCodeUnknown)
	}

	if err := s.event.Updated(ctx, old, ent); err != nil {
		return nil, NewError(err, err.Error(), ErrCodeUnknown)
	}

	return ent, nil
}

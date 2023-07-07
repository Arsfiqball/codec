package resource

import (
	"context"
	"feature/internal/value/domain"
	"feature/internal/value/user"
)

func (s *Service) GetList(ctx context.Context, query domain.Query, user user.Entity) ([]domain.Entity, error) {
	ctx, span := s.tracer.Start(ctx, "feature/internal/application/resource.Service.GetList")
	defer span.End()

	if err := query.Validate(); err != nil {
		return nil, NewError(err, err.Error(), ErrCodeInvalidQuery)
	}

	ents, err := s.repo.GetList(ctx, query)
	if err != nil {
		return nil, NewError(err, err.Error(), ErrCodeUnknown)
	}

	for _, ent := range ents {
		if err := ent.AuthorizeRead(user); err != nil {
			return nil, NewError(err, err.Error(), ErrCodeUnauthorized)
		}
	}

	return ents, nil
}

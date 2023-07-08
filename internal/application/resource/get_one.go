package resource

import (
	"context"
	"feature/internal/value/domain"
	"feature/internal/value/user"
)

func (s *Service) GetOne(ctx context.Context, query domain.Query, user user.Identity) (domain.Entity, error) {
	ctx, span := s.tracer.Start(ctx, "feature/internal/application/resource.Service.GetOne")
	defer span.End()

	if err := query.Validate(); err != nil {
		return domain.Entity{}, NewError(err, err.Error(), ErrCodeInvalidQuery)
	}

	ent, err := s.repo.GetOne(ctx, query)
	if err != nil {
		return domain.Entity{}, NewError(err, err.Error(), ErrCodeUnknown)
	}

	if err := s.authorizeGetOne(user, ent); err != nil {
		return domain.Entity{}, NewError(err, err.Error(), ErrCodeUnauthorized)
	}

	return ent, nil
}

func (s *Service) authorizeGetOne(u user.Identity, ent domain.Entity) error {
	// TODO: implement authorization logic

	return nil
}

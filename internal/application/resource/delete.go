package resource

import (
	"context"
	"errors"
	"feature/internal/value/domain"
	"feature/internal/value/user"
)

func (s *Service) Delete(ctx context.Context, query domain.Query, user user.Identity) (domain.Entity, error) {
	ctx, span := s.tracer.Start(ctx, "feature/internal/application/resource.Service.Delete")
	defer span.End()

	if err := query.Validate(); err != nil {
		return nil, NewError(err, err.Error(), ErrCodeInvalidQuery)
	}

	ent, err := s.repo.GetOne(ctx, query)
	if err != nil {
		return nil, NewError(err, err.Error(), ErrCodeUnknown)
	}

	if err := s.authorizeDelete(user, ent); err != nil {
		return nil, NewError(err, err.Error(), ErrCodeUnauthorized)
	}

	ent, err = s.repo.Delete(ctx, ent)
	if err != nil {
		return nil, NewError(err, err.Error(), ErrCodeUnknown)
	}

	if err := s.event.Deleted(ctx, ent); err != nil {
		return nil, NewError(err, err.Error(), ErrCodeUnknown)
	}

	return ent, nil
}

func (s *Service) authorizeDelete(u user.Identity, ent domain.Entity) error {
	// TODO: implement authorization logic

	if u.Provider() == user.ProviderGuest {
		return errors.New("guest user can't delete resource")
	}

	return nil
}

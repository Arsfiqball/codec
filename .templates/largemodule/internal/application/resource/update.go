package resource

import (
	"context"
	"errors"

	"github.com/Arsfiqball/codec/internal/value/domain"
	"github.com/Arsfiqball/codec/internal/value/user"
)

func (s *Service) Update(ctx context.Context, query domain.Query, patch domain.Patch, user user.Identity) (domain.Entity, error) {
	ctx, span := s.tracer.Start(ctx, "github.com/Arsfiqball/codec/internal/application/resource.Service.Update")
	defer span.End()

	if err := query.Validate(); err != nil {
		return domain.Entity{}, NewError(err, err.Error(), ErrCodeInvalidQuery)
	}

	old, err := s.repo.GetOne(ctx, query)
	if err != nil {
		return domain.Entity{}, NewError(err, err.Error(), ErrCodeUnknown)
	}

	ent := domain.NewEntityClone(old)

	if err := ent.Patch(patch); err != nil {
		return domain.Entity{}, NewError(err, err.Error(), ErrCodeInvalidEntity)
	}

	if err := ent.Validate(); err != nil {
		return domain.Entity{}, NewError(err, err.Error(), ErrCodeInvalidEntity)
	}

	if err := s.authorizeUpdate(user, ent); err != nil {
		return domain.Entity{}, NewError(err, err.Error(), ErrCodeUnauthorized)
	}

	ent, err = s.repo.Update(ctx, ent)
	if err != nil {
		return domain.Entity{}, NewError(err, err.Error(), ErrCodeUnknown)
	}

	if err := s.event.Updated(ctx, old, ent); err != nil {
		return domain.Entity{}, NewError(err, err.Error(), ErrCodeUnknown)
	}

	return ent, nil
}

func (s *Service) authorizeUpdate(u user.Identity, ent domain.Entity) error {
	// TODO: implement authorization logic

	if u.Provider() == user.ProviderGuest {
		return errors.New("guest user can't update resource")
	}

	return nil
}

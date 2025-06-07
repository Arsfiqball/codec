package resource

import (
	"context"

	"github.com/Arsfiqball/codec/internal/value/domain"
	"github.com/Arsfiqball/codec/internal/value/user"
)

func (s *Service) GetList(ctx context.Context, query domain.Query, user user.Identity) ([]domain.Entity, error) {
	ctx, span := s.tracer.Start(ctx, "github.com/Arsfiqball/codec/internal/application/resource.Service.GetList")
	defer span.End()

	if err := query.Validate(); err != nil {
		return nil, NewError(err, err.Error(), ErrCodeInvalidQuery)
	}

	ents, err := s.repo.GetList(ctx, query)
	if err != nil {
		return nil, NewError(err, err.Error(), ErrCodeUnknown)
	}

	if err := s.authorizeGetList(user, ents); err != nil {
		return nil, NewError(err, err.Error(), ErrCodeUnauthorized)
	}

	return ents, nil
}

func (s *Service) authorizeGetList(u user.Identity, ents []domain.Entity) error {
	// TODO: implement authorization logic

	return nil
}

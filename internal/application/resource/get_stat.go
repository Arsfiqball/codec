package resource

import (
	"context"
	"feature/internal/value/domain"
	"feature/internal/value/user"
)

func (s *Service) GetStat(ctx context.Context, query domain.Query, user user.Entity) ([]domain.Stat, error) {
	ctx, span := s.tracer.Start(ctx, "feature/internal/application/resource.Service.GetStat")
	defer span.End()

	if err := query.Validate(); err != nil {
		return nil, NewError(err, err.Error(), ErrCodeInvalidQuery)
	}

	stat, err := s.repo.GetStat(ctx, query)
	if err != nil {
		return nil, NewError(err, err.Error(), ErrCodeUnknown)
	}

	return stat, nil
}

package job

import (
	"context"
	"feature/internal/value/domain"
	"time"
)

func (s *Service) ResetEverybodyPassword(ctx context.Context, before time.Time) error {
	ctx, span := s.tracer.Start(ctx, "feature/internal/application/job.Service.ResetEverybodyPassword")
	defer span.End()

	query := domain.Query{}
	query.Limit = 1000

	ents, err := s.repo.GetList(ctx, query)
	if err != nil {
		return err
	}

	for _, ent := range ents {
		updated := domain.NewEntityClone(ent)
		updated.ResetPassword()

		_, err := s.repo.Update(ctx, updated)
		if err != nil {
			return err
		}
	}

	return nil
}

package scheduler

import (
	"context"
	"time"
)

func (s *Service) ResetEverybodyPassword(ctx context.Context) error {
	ctx, span := s.tracer.Start(ctx, "github.com/Arsfiqball/codec/internal/application/scheduler.Service.ResetEverybodyPassword")
	defer span.End()

	return s.queuer.ResetEverybodyPassword(ctx, time.Now().Add(-24*time.Hour))
}

package queue

import (
	"context"
	"time"
)

type Queuer interface {
	ResetEverybodyPassword(ctx context.Context, before time.Time) error
}

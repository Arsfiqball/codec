package something

import "context"

//go:generate mockery --name Repo --inpackage --case snake
type Repo interface {
	Aggregate(ctx context.Context, query Query) ([]Entity, error)
	AggregateStats(ctx context.Context, query Query) ([]Stat, error)
	Save(ctx context.Context, entity Entity) error
	GetByID(ctx context.Context, id string) (Entity, error)
	DeleteByID(ctx context.Context, id string) error
}

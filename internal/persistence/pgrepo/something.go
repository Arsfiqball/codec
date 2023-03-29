package pgrepo

import (
	"context"
	"database/sql"
	"feature/internal/value/something"
)

// Repo is the repository for something.
type Something struct {
	db *sql.DB
}

// NewSomething creates a new repository for something.
func NewSomething(db *sql.DB) *Something {
	return &Something{
		db: db,
	}
}

// Aggregate returns a list of something entities.
func (r *Something) Aggregate(ctx context.Context, query something.Query) ([]something.Entity, error) {
	// TODO: Implement this.

	return []something.Entity{}, nil
}

// AggregateStats returns a list of something stats.
func (r *Something) AggregateStats(ctx context.Context, query something.Query) ([]something.Stat, error) {
	// TODO: Implement this.

	return []something.Stat{}, nil
}

// Save saves a something entity.
func (r *Something) Save(ctx context.Context, entity something.Entity) error {
	// TODO: Implement this.

	return nil
}

// GetByID returns a something entity by ID.
func (r *Something) GetByID(ctx context.Context, id string) (something.Entity, error) {
	// TODO: Implement this.

	return something.Entity{}, nil
}

// DeleteByID deletes a list of something entities by ID.
func (r *Something) DeleteByID(ctx context.Context, id string) error {
	// TODO: Implement this.

	return nil
}

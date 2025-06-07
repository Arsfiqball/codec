package domain

import "context"

type Event interface {
	Created(ctx context.Context, ent Entity) error
	Updated(ctx context.Context, oldEnt Entity, newEnt Entity) error
	Deleted(ctx context.Context, ent Entity) error
}

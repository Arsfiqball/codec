package domain

import "context"

type Repo interface {
	Create(context.Context, Entity) (Entity, error)
	Update(context.Context, Entity) (Entity, error)
	Delete(context.Context, Entity) (Entity, error)
	GetOne(context.Context, Query) (Entity, error)
	GetList(context.Context, Query) ([]Entity, error)
	GetStat(context.Context, Query) ([]Stat, error)
}

package domain

import "github.com/google/uuid"

func NewEntity() Entity {
	return &entityState{
		id: uuid.NewString(),
	}
}

func NewEntityClone(ent Entity) Entity {
	return &entityState{
		id:       ent.ID(),
		name:     ent.Name(),
		email:    ent.Email(),
		password: ent.Password(),
	}
}

func NewEntityWithData(
	id string,
	name string,
	email string,
	password string,
) Entity {
	return &entityState{
		id:       id,
		name:     name,
		email:    email,
		password: password,
	}
}

func NewStat(
	name string,
	count int,
) Stat {
	return &statState{
		name:  name,
		count: count,
	}
}

func NewOmittable[T any](value T) Omittable[T] {
	return &omittableState[T]{
		value: value,
		valid: true,
	}
}

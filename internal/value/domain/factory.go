package domain

import "github.com/google/uuid"

func NewEntity() Entity {
	return Entity{
		id: uuid.NewString(),
	}
}

func NewEntityClone(ent Entity) Entity {
	return ent
}

func NewEntityWithData(
	id string,
	name string,
	email string,
	password string,
) Entity {
	return Entity{
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
	return Stat{
		name:  name,
		count: count,
	}
}

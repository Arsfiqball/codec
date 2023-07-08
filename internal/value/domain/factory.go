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

func NewQuery() Query {
	return &queryState{}
}

func NewQueryWithData(
	conditions []Condition,
	search string,
	limit int,
	skip int,
	sort []Sort,
	accumulator []string,
	group []string,
	with []string,
) Query {
	return &queryState{
		conditions:  conditions,
		search:      search,
		limit:       limit,
		skip:        skip,
		sort:        sort,
		accumulator: accumulator,
		group:       group,
		with:        with,
	}
}

func NewCondition(
	field string,
	op string,
	values []string,
) Condition {
	return &conditionState{
		field:  field,
		op:     op,
		values: values,
	}
}

func NewSort(
	field string,
	desc bool,
) Sort {
	return &sortState{
		field: field,
		desc:  desc,
	}
}

func NewStat(
	fields []string,
	count int,
) Stat {
	return &statState{
		fields: fields,
		count:  count,
	}
}

func NewOmittable[T any](value T) Omittable[T] {
	return &omittableState[T]{
		value: value,
		valid: true,
	}
}

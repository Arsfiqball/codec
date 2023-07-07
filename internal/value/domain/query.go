package domain

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

type Condition interface {
	Field() string
	Op() string
	Values() []string
	Validate() error
}

type conditionState struct {
	field  string
	op     string
	values []string
}

func (c conditionState) Field() string {
	return c.field
}

func (c conditionState) Op() string {
	return c.op
}

func (c conditionState) Values() []string {
	return c.values
}

func (c conditionState) Validate() error {
	errs := validation.Errors{
		"field":  validation.Validate(c.field, validation.Required, validation.In(stringsToInterfaces(QueryableFields)...)),
		"op":     validation.Validate(c.op, validation.Required),
		"values": validation.Validate(c.values, validation.Required),
	}

	return errs.Filter()
}

type Sort interface {
	Field() string
	Desc() bool
}

type sortState struct {
	field string
	desc  bool
}

func (s sortState) Field() string {
	return s.field
}

func (s sortState) Desc() bool {
	return s.desc
}

type Query interface {
	Conditions() []Condition
	Limit() int
	Skip() int
	Sort() []Sort
	Accumulator() []string
	Group() []string
	With() []string
	Validate() error
}

type queryState struct {
	conditions  []Condition
	limit       int
	skip        int
	sort        []Sort
	accumulator []string
	group       []string
	with        []string
}

func (q queryState) Conditions() []Condition {
	return q.conditions
}

func (q queryState) Limit() int {
	return q.limit
}

func (q queryState) Skip() int {
	return q.skip
}

func (q queryState) Sort() []Sort {
	return q.sort
}

func (q queryState) Accumulator() []string {
	return q.accumulator
}

func (q queryState) Group() []string {
	return q.group
}

func (q queryState) With() []string {
	return q.with
}

func (q queryState) Validate() error {
	errs := validation.Errors{}

	for _, c := range q.conditions {
		if err := c.Validate(); err != nil {
			errs["conditions"] = err
		}
	}

	errs["limit"] = validation.Validate(q.limit, validation.Min(1))
	errs["skip"] = validation.Validate(q.skip, validation.Min(0))
	errs["sort"] = validation.Validate(q.sort, validation.In(stringsToInterfaces(SortableFields)...))
	errs["accumulator"] = validation.Validate(q.accumulator, validation.In(stringsToInterfaces(AccumulableFields)...))
	errs["group"] = validation.Validate(q.group, validation.In(stringsToInterfaces(GroupableFields)...))
	errs["with"] = validation.Validate(q.with, validation.In(stringsToInterfaces(WithableFields)...))

	return errs.Filter()
}

func stringsToInterfaces(ss []string) []interface{} {
	is := make([]interface{}, len(ss))

	for i, v := range ss {
		is[i] = v
	}

	return is
}

package parser

import "encoding/json"

type omittable[T any] struct {
	value T
	valid bool
}

func (o omittable[T]) Value() T {
	return o.value
}

func (o omittable[T]) Valid() bool {
	return o.valid
}

func (o *omittable[T]) Unmarshal(b []byte) error {
	if err := json.Unmarshal(b, &o.value); err != nil {
		return err
	}

	o.valid = true

	return nil
}

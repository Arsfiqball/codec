package domain

type Omittable[T any] interface {
	Value() T
	Valid() bool
}

type omittableState[T any] struct {
	value T
	valid bool
}

func (o omittableState[T]) Value() T {
	return o.value
}

func (o omittableState[T]) Valid() bool {
	return o.valid
}

type Patch struct {
	Name     Omittable[string]
	Email    Omittable[string]
	Password Omittable[string]
}

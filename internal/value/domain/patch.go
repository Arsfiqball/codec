package domain

type Omittable[T any] interface {
	Value() T
	Valid() bool
}

type Patch struct {
	Name     Omittable[string]
	Email    Omittable[string]
	Password Omittable[string]
}

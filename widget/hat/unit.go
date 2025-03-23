package hat

import "encoding/json"

type Unit[T any] struct {
	present bool
	valid   bool
	value   T
}

func (u Unit[T]) IsPresent() bool {
	return u.present
}

func (u Unit[T]) IsNull() bool {
	return !u.valid
}

func (u Unit[T]) Value() T {
	return u.value
}

func (u Unit[T]) MarshalJSON() ([]byte, error) {
	if u.present {
		if u.valid {
			return json.Marshal(u.value)
		}

		return []byte("null"), nil
	}

	return []byte("{}"), nil // {} represents not present
}

func (u *Unit[T]) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		*u = Null[T]()
		return nil
	}

	*u = Value(u.value)
	return json.Unmarshal(data, &u.value)
}

func NotPresent[T any]() Unit[T] {
	return Unit[T]{}
}

func Null[T any]() Unit[T] {
	return Unit[T]{present: true}
}

func Value[T any](v T) Unit[T] {
	return Unit[T]{present: true, valid: true, value: v}
}

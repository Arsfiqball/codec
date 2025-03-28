package lord

import "sync"

type Collection[T any] []T

type Result[T any] struct {
	Value T
	Err   error
}

func WithResult[T any](t T) Result[T] {
	return Result[T]{Value: t}
}

func Range[T any](n int, gen func(int) T) Collection[T] {
	result := make(Collection[T], n)

	for i := 0; i < n; i++ {
		result[i] = gen(i)
	}

	return result
}

func Collect[T any](values ...T) Collection[T] {
	return values
}

func Convert[T, U any](c Collection[T], convert func(T) U) Collection[U] {
	result := make(Collection[U], len(c))

	for i, v := range c {
		result[i] = convert(v)
	}

	return result
}

func (c Collection[T]) Add(values ...T) Collection[T] {
	return append(c, values...)
}

func (c Collection[T]) Remove(match func(T) bool) Collection[T] {
	var result Collection[T]

	for _, v := range c {
		if !match(v) {
			result = append(result, v)
		}
	}

	return result
}

func (c Collection[T]) Find(match func(T) bool) (T, bool) {
	var zero T

	for _, v := range c {
		if match(v) {
			return v, true
		}
	}

	return zero, false
}

func (c Collection[T]) Filter(match func(T) bool) Collection[T] {
	var result Collection[T]

	for _, v := range c {
		if match(v) {
			result = append(result, v)
		}
	}

	return result
}

func (c Collection[T]) Map(mut func(T) T) Collection[T] {
	for i, v := range c {
		c[i] = mut(v)
	}

	return c
}

func (c Collection[T]) ParallelMap(mut func(T) T) Collection[T] {
	var wg sync.WaitGroup
	wg.Add(len(c))

	for i, v := range c {
		go func(i int, v T) {
			defer wg.Done()
			c[i] = mut(v)
		}(i, v)
	}

	wg.Wait()

	return c
}

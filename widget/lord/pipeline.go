package lord

import "context"

type Pipe[T any] func(context.Context, T) (T, error)

func Pipeline[T any](pipes ...Pipe[T]) Pipe[T] {
	return func(ctx context.Context, v T) (T, error) {
		for _, pipe := range pipes {
			var err error
			v, err = pipe(ctx, v)
			if err != nil {
				return v, err
			}
		}

		return v, nil
	}
}

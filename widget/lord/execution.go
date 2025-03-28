package lord

import (
	"context"
	"errors"
	"sync"
	"time"
)

type Executor[T any] func(context.Context, T) error

func Sequential[T any](exe ...Executor[T]) Executor[T] {
	return func(ctx context.Context, v T) error {
		for _, callback := range exe {
			if err := callback(ctx, v); err != nil {
				return err
			}
		}

		return nil
	}
}

func Parallel[T any](exe ...Executor[T]) Executor[T] {
	return func(ctx context.Context, v T) error {
		var wg sync.WaitGroup

		errChan := make(chan error, len(exe))

		for _, callback := range exe {
			wg.Add(1)

			go func(w *sync.WaitGroup, callback Executor[T]) {
				defer w.Done()
				errChan <- callback(ctx, v)
			}(&wg, callback)
		}

		wg.Wait()
		errs := []error{}

		for i := 0; i < len(exe); i++ {
			if err := <-errChan; err != nil {
				errs = append(errs, err)
			}
		}

		return errors.Join(errs...)
	}
}

func Timeout[T any](exe Executor[T], timeout time.Duration) Executor[T] {
	return func(ctx context.Context, v T) error {
		ctx, cancel := context.WithTimeout(ctx, timeout)
		defer cancel()

		return exe(ctx, v)
	}
}

func Retry[T any](exe Executor[T], retries int, delay time.Duration) Executor[T] {
	return func(ctx context.Context, v T) error {
		var err error

		for i := 0; i < retries; i++ {
			err = exe(ctx, v)
			if err == nil {
				return nil
			}

			time.Sleep(delay)
		}

		return err
	}
}

func IgnoreError[T any](exe Executor[T]) Executor[T] {
	return func(ctx context.Context, v T) error {
		_ = exe(ctx, v)
		return nil
	}
}

func Atomic[T any](exe Executor[T], rollback Executor[T]) Executor[T] {
	return func(ctx context.Context, v T) error {
		err := exe(ctx, v)
		if err != nil {
			return rollback(ctx, v)
		}

		return nil
	}
}

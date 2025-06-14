package skipper

import (
	"context"
	"errors"
)

var (
	ErrNotFound = errors.New("not found")
)

type Client interface {
	RequestSomething(context.Context, RequestSomething) (Something, error)
}

type RequestSomething struct {
	// TODO: Add your input fields here
}

package scaper

import (
	"context"
	"errors"
)

var (
	ErrNotFound = errors.New("not found")
)

type Database interface {
	GetSomething(context.Context, GetSomething) (Something, error)
}

type GetSomething struct {
	// TODO: Add your input fields here
}

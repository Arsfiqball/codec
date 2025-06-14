//go:build wireinject
// +build wireinject

package module

import (
	"context"

	"github.com/google/wire"
)

func New(ctx context.Context, cfg Config) (*Module, error) {
	wire.Build(RegisterSet)
	return &Module{}, nil
}

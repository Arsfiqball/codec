//go:build wireinject
// +build wireinject

package featurepkg

import (
	"context"

	"github.com/google/wire"
)

func New(ctx context.Context, cfg Config) (*FeatureName, error) {
	wire.Build(RegisterSet)
	return &FeatureName{}, nil
}

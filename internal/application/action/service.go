package action

import (
	"context"
	"feature/internal/value/something"
)

// Service is the application service for resource.
type IService interface {
	ClaimSomething(ctx context.Context, input ClaimSomethingInput) (ClaimSomethingOutput, error)
}

// Service is the application service for resource.
type Service struct {
	somethingRepo something.Repo
}

// NewService creates a new resource service.
func NewService(somethingRepo something.Repo) *Service {
	return &Service{
		somethingRepo: somethingRepo,
	}
}

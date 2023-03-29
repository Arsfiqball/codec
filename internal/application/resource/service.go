package resource

import (
	"context"
	"feature/internal/value/something"
)

// Service is the application service for resource.
type IService interface {
	GetList(ctx context.Context, input GetListInput) (GetListOutput, error)
	GetOne(ctx context.Context, input GetOneInput) (GetOneOutput, error)
	GetStat(ctx context.Context, input GetStatInput) (GetStatOutput, error)
	Create(ctx context.Context, input CreateInput) (CreateOutput, error)
	Update(ctx context.Context, input UpdateInput) (UpdateOutput, error)
	Delete(ctx context.Context, input DeleteInput) (DeleteOutput, error)
	BulkOps(ctx context.Context, input BulkOpsInput) (BulkOpsOutput, error)
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

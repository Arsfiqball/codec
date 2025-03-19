package decoder

import (
	"context"
	"errors"

	"go.opentelemetry.io/otel/trace"
)

type Service interface {
	GetList(ctx context.Context, input GetListInput) (GetListOutput, error)
	GetOne(ctx context.Context, input GetOneInput) (GetOneOutput, error)
	GetStat(ctx context.Context, input GetStatInput) (GetStatOutput, error)
	PostOne(ctx context.Context, input PostOneInput) (PostOneOutput, error)
	PatchOne(ctx context.Context, input PatchOneInput) (PatchOneOutput, error)
	DeleteOne(ctx context.Context, input DeleteOneInput) (DeleteOneOutput, error)
	BulkOps(ctx context.Context, input BulkOpsInput) (BulkOpsOutput, error)
}

var _ Service = (*service)(nil)

type service struct {
	tracer trace.Tracer
}

func NewService(
	tracer trace.Tracer,
) (Service, error) {
	if tracer == nil {
		return nil, errors.New("tracer is required")
	}

	return &service{
		tracer: tracer,
	}, nil
}

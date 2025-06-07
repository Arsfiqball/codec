package resource

import (
	"context"

	"github.com/Arsfiqball/codec/internal/value/domain"
	"github.com/Arsfiqball/codec/internal/value/user"
)

type Ops struct {
	Type  string
	Query domain.Query
	Patch domain.Patch
}

type OpsResult struct {
	Type   string
	Entity domain.Entity
	Error  error
}

func (s *Service) BulkOps(ctx context.Context, ops []Ops, user user.Identity) ([]OpsResult, error) {
	ctx, span := s.tracer.Start(ctx, "github.com/Arsfiqball/codec/internal/application/resource.Service.BulkOps")
	defer span.End()

	var results []OpsResult

	for _, op := range ops {
		switch op.Type {
		case "create":
			ent, err := s.Create(ctx, op.Patch, user)
			results = append(results, OpsResult{
				Type:   op.Type,
				Entity: ent,
				Error:  err,
			})
		case "update":
			ent, err := s.Update(ctx, op.Query, op.Patch, user)
			results = append(results, OpsResult{
				Type:   op.Type,
				Entity: ent,
				Error:  err,
			})
		case "delete":
			ent, err := s.Delete(ctx, op.Query, user)
			results = append(results, OpsResult{
				Type:   op.Type,
				Entity: ent,
				Error:  err,
			})
		default:
			results = append(results, OpsResult{
				Type:   op.Type,
				Entity: domain.Entity{},
				Error:  NewError(nil, "invalid operation type", ErrCodeInvalidOpsType),
			})
		}
	}

	return results, nil
}

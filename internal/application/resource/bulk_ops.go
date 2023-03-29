package resource

import (
	"context"
	"feature/internal/value/something"
)

// Ops is the operations for the BulkOps method.
type Ops struct {
	Type  string          `json:"type"`
	Query something.Query `json:"query"`
	Patch something.Patch `json:"patch"`
}

// BulkOpsInput is the input for the BulkOps method.
type BulkOpsInput struct {
	Ops []Ops `json:"ops"`
}

// BulkOpsOutput is the output for the BulkOps method.
type BulkOpsOutput struct {
	Docs []something.Entity `json:"docs"`
}

// BulkOps performs bulk operations.
func (s *Service) BulkOps(ctx context.Context, input BulkOpsInput) (BulkOpsOutput, error) {
	docs := []something.Entity{}

	for _, ops := range input.Ops {
		switch ops.Type {
		case "create":
			output, err := s.Create(ctx, CreateInput{Patch: ops.Patch})
			if err != nil {
				return BulkOpsOutput{}, err
			}

			docs = append(docs, output.Doc)
		case "update":
			output, err := s.Update(ctx, UpdateInput{Query: ops.Query, Patch: ops.Patch})
			if err != nil {
				return BulkOpsOutput{}, err
			}

			docs = append(docs, output.Doc)
		case "delete":
			output, err := s.Delete(ctx, DeleteInput{Query: ops.Query})
			if err != nil {
				return BulkOpsOutput{}, err
			}

			docs = append(docs, output.Doc)
		}
	}

	return BulkOpsOutput{Docs: docs}, nil
}

package resource

import "context"

type BulkOpsInput struct {
	// TODO: Add your input fields here
}

type BulkOpsOutput struct {
	// TODO: Add your output fields here
}

func (s *service) BulkOps(ctx context.Context, input BulkOpsInput) (BulkOpsOutput, error) {
	ctx, span := s.tracer.Start(ctx, "gadget/internal/resource.Service.BulkOps") // Don't forget to change the span name
	defer span.End()

	// TODO: Implement your business logic here

	return BulkOpsOutput{}, nil
}

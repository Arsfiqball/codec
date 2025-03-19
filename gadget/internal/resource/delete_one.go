package resource

import "context"

type DeleteOneInput struct {
	// TODO: Add your input fields here
}

type DeleteOneOutput struct {
	// TODO: Add your output fields here
}

func (s *service) DeleteOne(ctx context.Context, input DeleteOneInput) (DeleteOneOutput, error) {
	ctx, span := s.tracer.Start(ctx, "gadget/internal/resource.Service.DeleteOne") // Don't forget to change the span name
	defer span.End()

	// TODO: Implement your business logic here

	return DeleteOneOutput{}, nil
}

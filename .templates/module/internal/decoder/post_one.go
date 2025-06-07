package decoder

import "context"

type PostOneInput struct {
	// TODO: Add your input fields here
}

type PostOneOutput struct {
	// TODO: Add your output fields here
}

func (s *service) PostOne(ctx context.Context, input PostOneInput) (PostOneOutput, error) {
	ctx, span := s.tracer.Start(ctx, "gadget/internal/resource.Service.PostOne") // Don't forget to change the span name
	defer span.End()

	// TODO: Implement your business logic here

	return PostOneOutput{}, nil
}

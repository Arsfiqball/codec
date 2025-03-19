package decoder

import "context"

type GetOneInput struct {
	// TODO: Add your input fields here
}

type GetOneOutput struct {
	// TODO: Add your output fields here
}

func (s *service) GetOne(ctx context.Context, input GetOneInput) (GetOneOutput, error) {
	ctx, span := s.tracer.Start(ctx, "gadget/internal/resource.Service.GetOne") // Don't forget to change the span name
	defer span.End()

	// TODO: Implement your business logic here

	return GetOneOutput{}, nil
}

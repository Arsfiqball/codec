package decoder

import "context"

type PatchOneInput struct {
	// TODO: Add your input fields here
}

type PatchOneOutput struct {
	// TODO: Add your output fields here
}

func (s *service) PatchOne(ctx context.Context, input PatchOneInput) (PatchOneOutput, error) {
	ctx, span := s.tracer.Start(ctx, "gadget/internal/resource.Service.PatchOne") // Don't forget to change the span name
	defer span.End()

	// TODO: Implement your business logic here

	return PatchOneOutput{}, nil
}

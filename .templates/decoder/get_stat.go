package decoder

import "context"

type GetStatInput struct {
	// TODO: Add your input fields here
}

type GetStatOutput struct {
	// TODO: Add your output fields here
}

func (s *service) GetStat(ctx context.Context, input GetStatInput) (GetStatOutput, error) {
	ctx, span := s.tracer.Start(ctx, "module/internal/resource.Service.GetStat") // Don't forget to change the span name
	defer span.End()

	// TODO: Implement your business logic here

	return GetStatOutput{}, nil
}

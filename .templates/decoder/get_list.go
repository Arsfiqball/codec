package decoder

import "context"

type GetListInput struct {
	// TODO: Add your input fields here
}

type GetListOutput struct {
	// TODO: Add your output fields here
}

func (s *service) GetList(ctx context.Context, input GetListInput) (GetListOutput, error) {
	ctx, span := s.tracer.Start(ctx, "module/internal/resource.Service.GetList") // Don't forget to change the span name
	defer span.End()

	// TODO: Implement your business logic here

	return GetListOutput{}, nil
}

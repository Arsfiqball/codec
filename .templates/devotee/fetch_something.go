package devotee

import "context"

type FetchSomethingInput struct {
	// TODO: Add your input fields here
}

type FetchSomethingOutput struct {
	// TODO: Add your output fields here
}

func (s *service) FetchSomething(ctx context.Context, input FetchSomethingInput) (FetchSomethingOutput, error) {
	ctx, span := s.tracer.Start(ctx, "module/internal/content.Service.FetchSomething")
	defer span.End()

	// TODO: Implement your business logic here

	return FetchSomethingOutput{}, nil
}

package disciple

import "context"

type DoSomethingInput struct {
	// TODO: Add your input fields here
}

type DoSomethingOutput struct {
	// TODO: Add your output fields here
}

func (s *service) DoSomething(ctx context.Context, input DoSomethingInput) (DoSomethingOutput, error) {
	ctx, span := s.tracer.Start(ctx, "module/internal/action.Service.DoSomething")
	defer span.End()

	// TODO: Implement your business logic here

	return DoSomethingOutput{}, nil
}

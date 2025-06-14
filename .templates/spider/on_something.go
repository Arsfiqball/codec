package spider

import "context"

type OnSomethingInput struct {
	// TODO: Add your input fields here
}

type OnSomethingOutput struct {
	// TODO: Add your output fields here
}

func (s *service) OnSomething(ctx context.Context, input OnSomethingInput) (OnSomethingOutput, error) {
	ctx, span := s.tracer.Start(ctx, "module/internal/listener.Service.OnSomething")
	defer span.End()

	// TODO: Implement your business logic here

	return OnSomethingOutput{}, nil
}

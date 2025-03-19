package bumper

import "context"

type ScheduleSomethingInput struct {
	// TODO: Add your input fields here
}

type ScheduleSomethingOutput struct {
	// TODO: Add your output fields here
}

func (s *service) ScheduleSomething(ctx context.Context, input ScheduleSomethingInput) (ScheduleSomethingOutput, error) {
	ctx, span := s.tracer.Start(ctx, "gadget/internal/listener.Service.OnSomething")
	defer span.End()

	// TODO: Implement your business logic here

	return ScheduleSomethingOutput{}, nil
}

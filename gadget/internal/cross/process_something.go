package cross

import "context"

type ProcessSomethingInput struct {
	// TODO: Add your input fields here
}

type ProcessSomethingOutput struct {
	// TODO: Add your output fields here
}

func (s *service) ProcessSomething(ctx context.Context, input ProcessSomethingInput) (ProcessSomethingOutput, error) {
	ctx, span := s.tracer.Start(ctx, "gadget/internal/job.Service.ProcessSomething")
	defer span.End()

	// TODO: Implement your business logic here

	return ProcessSomethingOutput{}, nil
}

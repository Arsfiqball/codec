package bumper

import (
	"context"
)

type ScheduleSomething struct {
	Params struct {
		ID string `param:"id"`
	}

	Query struct {
		// TODO: Add your query fields here
	}

	Body struct {
		// TODO: Add your body fields here
	}
}

type Something struct {
	// TODO: Add your output fields here
}

func (s *service) ScheduleSomething(ctx context.Context, input ScheduleSomething) (Something, error) {
	_, span := s.tracer.Start(ctx, "gadget/internal/listener.Service.ScheduleSomething")
	defer span.End()

	// TODO: Implement your business logic here

	return Something{}, nil
}

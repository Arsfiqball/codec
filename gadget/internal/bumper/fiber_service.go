package bumper

import (
	"errors"
	"feature/widget/flame"

	"github.com/gofiber/fiber/v2"
	"go.opentelemetry.io/otel/trace"
)

type FiberService struct {
	tracer  trace.Tracer
	service Service
}

func NewFiberService(tracer trace.Tracer, service Service) (*FiberService, error) {
	if tracer == nil {
		return nil, errors.New("tracer is required")
	}

	if service == nil {
		return nil, errors.New("service is required")
	}

	return &FiberService{
		tracer:  tracer,
		service: service,
	}, nil
}

func (s *FiberService) ScheduleSomething(c *fiber.Ctx) error {
	ctx, span := s.tracer.Start(c.UserContext(), "gadget/internal/bumper.FiberService.ScheduleSomething")
	defer span.End()

	var input ScheduleSomething

	if err := c.ParamsParser(&input.Params); err != nil {
		return flame.BadRequest.Wrap(err).WithInfo("failed to parse params")
	}

	if err := c.QueryParser(&input.Query); err != nil {
		return flame.BadRequest.Wrap(err).WithInfo("failed to parse query")
	}

	if err := c.BodyParser(&input.Body); err != nil {
		return flame.BadRequest.Wrap(err).WithInfo("failed to parse body")
	}

	output, err := s.service.ScheduleSomething(ctx, input)
	if err != nil {
		return flame.Unexpected(err)
	}

	return c.JSON(output)
}

package reviver

import (
	"context"
	"errors"

	"github.com/ThreeDotsLabs/watermill/message"
	"go.opentelemetry.io/otel/trace"
)

var _ Event = &eventWatermill{}

type eventWatermill struct {
	tracer             trace.Tracer
	watermillPublisher message.Publisher
}

func NewEventWatermill(tracer trace.Tracer, watermillPublisher message.Publisher) (Event, error) {
	if tracer == nil {
		return nil, errors.New("tracer is required")
	}

	if watermillPublisher == nil {
		return nil, errors.New("watermillPublisher is required")
	}

	return &eventWatermill{
		tracer:             tracer,
		watermillPublisher: watermillPublisher,
	}, nil
}

func (e *eventWatermill) SomethingHappened(ctx context.Context, somethingHappened SomethingHappened) error {
	ctx, span := e.tracer.Start(ctx, "gadget/internal/reviver.Event.SomethingHappened") // Don't forget to change the span name
	defer span.End()

	// TODO: Implement your watermill event emitter here

	return nil
}

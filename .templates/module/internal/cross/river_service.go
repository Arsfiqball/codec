package cross

import (
	"context"
	"errors"

	"github.com/Arsfiqball/codec/widget/flame"

	"github.com/riverqueue/river"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

type RiverService struct {
	ProcessSomething *RiverProcessSomethingWorker
}

func NewRiverService(tracer trace.Tracer, service Service) (*RiverService, error) {
	if tracer == nil {
		return nil, errors.New("tracer is required")
	}

	if service == nil {
		return nil, errors.New("service is required")
	}

	return &RiverService{
		ProcessSomething: &RiverProcessSomethingWorker{
			tracer:  tracer,
			service: service,
		},
	}, nil
}

type RiverProcessSomethingWorker struct {
	river.WorkerDefaults[ProcessSomething]

	tracer  trace.Tracer
	service Service
}

func (w *RiverProcessSomethingWorker) Work(ctx context.Context, job *river.Job[ProcessSomething]) error {
	attrs := trace.WithAttributes(attribute.Int64("job_id", job.ID))

	ctx, span := w.tracer.Start(ctx, "gadget/internal/cross.RiverService.ProcessSomething", attrs)
	defer span.End()

	output, err := w.service.ProcessSomething(ctx, job.Args)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "failed to process job")

		return flame.Unexpected(err)
	}

	span.AddEvent("job completed", trace.WithAttributes(
		attribute.String("completed_at", output.CompletedAt.String()),
	))

	return nil
}

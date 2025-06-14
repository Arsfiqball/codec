package skipper

import (
	"context"
	"errors"
	"net/http"

	"go.opentelemetry.io/otel/trace"
)

var _ Client = (*clientHttp)(nil)

type clientHttp struct {
	tracer     trace.Tracer
	httpClient *http.Client
}

func NewClientHTTP(tracer trace.Tracer, httpClient *http.Client) (Client, error) {
	if tracer == nil {
		return nil, errors.New("tracer is required")
	}

	if httpClient == nil {
		return nil, errors.New("httpClient is required")
	}

	return &clientHttp{
		tracer:     tracer,
		httpClient: httpClient,
	}, nil
}

func (c *clientHttp) RequestSomething(ctx context.Context, input RequestSomething) (Something, error) {
	ctx, span := c.tracer.Start(ctx, "module/internal/skipper.Client.RequestSomething") // Don't forget to change the span name
	defer span.End()

	// TODO: Add your http client code here

	return Something{}, nil
}

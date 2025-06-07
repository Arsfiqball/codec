package scaper

import (
	"context"
	"errors"

	"go.opentelemetry.io/otel/trace"
	"gorm.io/gorm"
)

var _ Database = (*databaseGorm)(nil)

type databaseGorm struct {
	tracer trace.Tracer
	gormDB *gorm.DB
}

func NewClientHTTP(tracer trace.Tracer, gormDB *gorm.DB) (Database, error) {
	if tracer == nil {
		return nil, errors.New("tracer is required")
	}

	if gormDB == nil {
		return nil, errors.New("gormDB is required")
	}

	return &databaseGorm{
		tracer: tracer,
		gormDB: gormDB,
	}, nil
}

func (d *databaseGorm) GetSomething(ctx context.Context, input GetSomething) (Something, error) {
	ctx, span := d.tracer.Start(ctx, "gadget/internal/scaper.Database.GetSomething") // Don't forget to change the span name
	defer span.End()

	// TODO: Add your gorm code here

	return Something{}, nil
}

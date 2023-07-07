package featurepkg

import (
	"database/sql"
	"feature/internal/protocol/fiberhandler"

	"github.com/gofiber/fiber/v2"
	"github.com/google/wire"
	"go.opentelemetry.io/otel/trace"
)

// Config is the configuration for the feature.
type Config struct {
	Tracer   trace.Tracer
	Database *sql.DB
}

// FeatureName is the feature.
type FeatureName struct {
}

// Route routes the feature.
func (e *FeatureName) Route(router fiber.Router) {
	router.Get("/enum/sortable", fiberhandler.SortableFields)
	router.Get("/enum/queryable", fiberhandler.QueryableFields)
	router.Get("/enum/groupable", fiberhandler.GroupableFields)
	router.Get("/enum/withable", fiberhandler.WithableFields)
	router.Get("/enum/accumulable", fiberhandler.AccumulableFields)
}

// List of dependencies for the feature.
var RegisterSet = wire.NewSet(
	wire.Struct(new(FeatureName), "*"),
	wire.FieldsOf(new(Config), "Tracer", "Database"),
)

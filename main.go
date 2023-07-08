package featurepkg

import (
	"feature/internal/application/resource"
	"feature/internal/persistence/gormrepo"
	"feature/internal/persistence/wmpublisher"
	"feature/internal/protocol/fiberhandler"
	"feature/internal/value/domain"

	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/gofiber/fiber/v2"
	"github.com/google/wire"
	"go.opentelemetry.io/otel/trace"
	"gorm.io/gorm"
)

// Config is the configuration for the feature.
type Config struct {
	Tracer    trace.Tracer
	Database  *gorm.DB
	Publisher message.Publisher
}

// FeatureName is the feature.
type FeatureName struct {
	resourceHandler *fiberhandler.Resource
}

// Route routes the feature.
func (e *FeatureName) Route(router fiber.Router) {
	router.Get("/enum/sortable", fiberhandler.SortableFields)
	router.Get("/enum/queryable", fiberhandler.QueryableFields)
	router.Get("/enum/groupable", fiberhandler.GroupableFields)
	router.Get("/enum/withable", fiberhandler.WithableFields)
	router.Get("/enum/accumulable", fiberhandler.AccumulableFields)

	router.Post("/one", e.resourceHandler.Create)
	router.Patch("/one", e.resourceHandler.Update)
	router.Delete("/one", e.resourceHandler.Delete)
	router.Get("/one", e.resourceHandler.GetOne)
	router.Get("/list", e.resourceHandler.GetList)
	router.Get("/stat", e.resourceHandler.GetStat)
}

// List of dependencies for the feature.
var RegisterSet = wire.NewSet(
	fiberhandler.NewResource,
	resource.NewService,
	wire.Bind(new(resource.IService), new(*resource.Service)),
	gormrepo.NewDomain,
	wire.Bind(new(domain.Repo), new(*gormrepo.Domain)),
	wmpublisher.NewDomain,
	wire.Bind(new(domain.Event), new(*wmpublisher.Domain)),
	wire.Struct(new(FeatureName), "*"),
	wire.FieldsOf(new(Config), "Tracer", "Database", "Publisher"),
)

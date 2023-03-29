package featurepkg

import (
	"database/sql"
	"feature/internal/application/action"
	"feature/internal/application/resource"
	"feature/internal/persistence/pgrepo"
	"feature/internal/protocol/fiberhandler"
	"feature/internal/value/something"

	"github.com/gofiber/fiber/v2"
	"github.com/google/wire"
)

// Config is the configuration for the feature.
type Config struct {
	Database *sql.DB
}

// FeatureName is the feature.
type FeatureName struct {
	handler *fiberhandler.Handler
}

// Route routes the feature.
func (e *FeatureName) Route(router fiber.Router) {
	router.Post("/action/claim-something", e.handler.ClaimSomething)
	router.Post("/one", e.handler.Create)
	router.Patch("/one", e.handler.Update)
	router.Delete("/one", e.handler.Delete)
	router.Get("/one", e.handler.GetOne)
	router.Get("/list", e.handler.GetList)
	router.Get("/stat", e.handler.GetStat)
	router.Post("/bulk-ops", e.handler.BulkOps)
}

// List of dependencies for the feature.
var RegisterSet = wire.NewSet(
	pgrepo.NewSomething,
	action.NewService,
	resource.NewService,
	fiberhandler.NewHandler,
	wire.Bind(new(something.Repo), new(*pgrepo.Something)),
	wire.Bind(new(action.IService), new(*action.Service)),
	wire.Bind(new(resource.IService), new(*resource.Service)),
	wire.Struct(new(FeatureName), "*"),
	wire.FieldsOf(new(Config), "Database"),
)

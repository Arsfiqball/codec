package module

import (
	"github.com/google/wire"
	"go.opentelemetry.io/otel/sdk/trace"
	"gorm.io/gorm"
)

type Config struct {
	DB     *gorm.DB
	Tracer trace.Tracer
}

type Module struct {
}

var RegisterSet = wire.NewSet(
	wire.Struct(new(Module), "*"),
	wire.FieldsOf(new(Config), "Database", "Tracer"),
)

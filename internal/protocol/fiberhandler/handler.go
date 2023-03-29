package fiberhandler

import (
	"feature/internal/application/action"
	"feature/internal/application/resource"
)

// Handler is the handler for the fiber server.
type Handler struct {
	actionSvc   action.IService
	resourceSvc resource.IService
}

// NewHandler creates a new handler.
func NewHandler(actionSvc action.IService, resourceSvc resource.IService) *Handler {
	return &Handler{
		actionSvc:   actionSvc,
		resourceSvc: resourceSvc,
	}
}

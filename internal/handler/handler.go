package handler

import (
	"github.com/solumD/go-quotes-server/internal/service"
)

// handler is a struct that contains the service and handler functions.
type handler struct {
	service service.Service
}

// New returns a new handler.
func New(service service.Service) *handler {
	return &handler{
		service: service,
	}
}

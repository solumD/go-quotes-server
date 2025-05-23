package handler

import (
	"errors"

	"github.com/solumD/go-quotes-server/internal/service"
)

var (
	// ErrUnmarshalRequest is an error that is returned when the request is not unmarshaled.
	ErrUnmarshalRequest = errors.New("failed to unmarshal request")
	// ErrMarshalResponse is an error that is returned when the response is not marshaled.
	ErrMarshalResponse = errors.New("failed to marshal response")
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

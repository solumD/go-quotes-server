package handler

import (
	"encoding/json"

	"github.com/solumD/go-quotes-server/internal/service"
)

type handler struct {
	service service.Service
	decoder json.Decoder
}

func New(service service.Service) *handler {
	return &handler{
		service: service,
	}
}



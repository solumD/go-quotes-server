package handler

import "github.com/solumD/go-quotes-server/internal/service"

type handler struct {
	service service.Service
}

func New(service service.Service) *handler {
	return &handler{
		service: service,
	}
}

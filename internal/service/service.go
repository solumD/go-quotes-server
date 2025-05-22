package service

import (
	"context"
	"github.com/solumD/go-quotes-server/internal/model"
	"github.com/solumD/go-quotes-server/internal/repository"
)

type service struct {
	repo repository.Repository
}

func New(repository.Repository) Service {
	return &service{}
}
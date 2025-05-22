package repository

import (
	"context"

	"github.com/solumD/go-quotes-server/internal/model"
)

type Repository interface {
	SaveQuote(ctx context.Context, quote string, author string) (int64, error)
	GetAllQuotes(ctx context.Context) ([]*model.Quote, error)
	GetRandomQuote(ctx context.Context) (*model.Quote, error)
	GetQuotesByAuthor(ctx context.Context, author string) ([]*model.Quote, error)
	DeleteQuote(ctx context.Context, id int64) error
}

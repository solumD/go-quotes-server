package service

import (
	"context"

	"github.com/solumD/go-quotes-server/internal/model"
)

type Service interface {
	SaveQuote(ctx context.Context, quoteText string, quoteAuthor string) (int64, error)
	GetAllQuotes(ctx context.Context) ([]*model.Quote, error)
	GetRandomQuote(ctx context.Context) (*model.Quote, error)
	GetQuotesByAuthor(ctx context.Context, quoteAuthor string) ([]*model.Quote, error)
	DeleteQuote(ctx context.Context, quoteID int64) error
}

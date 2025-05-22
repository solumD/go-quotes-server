package repository

import (
	"context"

	"github.com/solumD/go-quotes-server/internal/model"
)

// Repository is an interface that defines the methods that a repository must implement.
type Repository interface {
	SaveQuote(ctx context.Context, quoteText string, quoteAuthor string) (int64, error)
	GetAllQuotes(ctx context.Context) ([]*model.Quote, error)
	GetRandomQuote(ctx context.Context) (*model.Quote, error)
	GetQuotesByAuthor(ctx context.Context, quoteAuthor string) ([]*model.Quote, error)
	DeleteQuote(ctx context.Context, ID int64) error
	Close()
}

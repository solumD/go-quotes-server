package srv

import (
	"context"
	"fmt"

	"github.com/solumD/go-quotes-server/internal/model"
	"github.com/solumD/go-quotes-server/internal/repository"
	"github.com/solumD/go-quotes-server/internal/service"
)

// srv is a struct that contains the repository.
type srv struct {
	repo repository.Repository
}

// New returns a new service.
func New(repo repository.Repository) service.Service {
	return &srv{
		repo: repo,
	}
}

// SaveQuote validates quote data and saves it in the database.
func (s *srv) SaveQuote(ctx context.Context, quoteText string, quoteAuthor string) (int64, error) {
	var fn = "srv.SaveQuote"

	if len(quoteText) == 0 {
		return 0, fmt.Errorf("%s: quote text is empty", fn)
	}

	if len(quoteAuthor) == 0 {
		return 0, fmt.Errorf("%s: quote author is empty", fn)
	}

	return s.repo.SaveQuote(ctx, quoteText, quoteAuthor)
}

// GetAllQuotes returns all quotes from the database.
func (s *srv) GetAllQuotes(ctx context.Context) ([]*model.Quote, error) {
	return s.repo.GetAllQuotes(ctx)
}

// GetRandomQuote returns a random quote from the database.
func (s *srv) GetRandomQuote(ctx context.Context) (*model.Quote, error) {
	return s.repo.GetRandomQuote(ctx)
}

// GetQuotesByAuthor validates quote author and returns quotes by this author from the database.
func (s *srv) GetQuotesByAuthor(ctx context.Context, quoteAuthor string) ([]*model.Quote, error) {
	var fn = "srv.GetQuotesByAuthor"
	if len(quoteAuthor) == 0 {
		return nil, fmt.Errorf("%s: quote author is empty", fn)
	}

	return s.repo.GetQuotesByAuthor(ctx, quoteAuthor)
}

// DeleteQuote deletes a quote from the database.
func (s *srv) DeleteQuote(ctx context.Context, id int64) error {
	var fn = "srv.DeleteQuote"
	if id <= 0 {
		return fmt.Errorf("%s: quote id is invalid", fn)
	}

	return s.repo.DeleteQuote(ctx, id)
}

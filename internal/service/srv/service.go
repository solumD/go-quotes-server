package srv

import (
	"context"

	"github.com/solumD/go-quotes-server/internal/model"
	"github.com/solumD/go-quotes-server/internal/repository"
	"github.com/solumD/go-quotes-server/internal/service"
	srverrors "github.com/solumD/go-quotes-server/internal/service/srv_errors"
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
	if len(quoteText) == 0 {
		return 0, srverrors.ErrTextIsEmpty
	}

	if len(quoteAuthor) == 0 {
		return 0, srverrors.ErrAuthorIsEmpty
	}

	id, err := s.repo.SaveQuote(ctx, quoteText, quoteAuthor)
	if err != nil {
		return 0, err
	}

	return id, nil
}

// GetAllQuotes returns all quotes from the database.
func (s *srv) GetAllQuotes(ctx context.Context) ([]*model.Quote, error) {
	return s.repo.GetAllQuotes(ctx)
}

// GetRandomQuote returns a random quote from the database.
func (s *srv) GetRandomQuote(ctx context.Context) (*model.Quote, error) {
	quote, err := s.repo.GetRandomQuote(ctx)
	if err != nil {
		return nil, err
	}

	return quote, nil
}

// GetQuotesByAuthor validates quote author and returns quotes by this author from the database.
func (s *srv) GetQuotesByAuthor(ctx context.Context, quoteAuthor string) ([]*model.Quote, error) {
	if len(quoteAuthor) == 0 {
		return nil, srverrors.ErrAuthorIsEmpty
	}

	quotes, err := s.repo.GetQuotesByAuthor(ctx, quoteAuthor)
	if err != nil {
		return nil, err
	}

	return quotes, nil
}

// DeleteQuote deletes a quote from the database.
func (s *srv) DeleteQuote(ctx context.Context, id int64) error {
	if id <= 0 {
		return srverrors.ErrInvalidID
	}

	err := s.repo.DeleteQuote(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

package srv

import (
	"context"
	"fmt"

	"github.com/solumD/go-quotes-server/internal/model"
	"github.com/solumD/go-quotes-server/internal/repository"
	"github.com/solumD/go-quotes-server/internal/service"
)

type srv struct {
	repo repository.Repository
}

func New(repo repository.Repository) service.Service {
	return &srv{
		repo: repo,
	}
}

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

func (s *srv) GetAllQuotes(ctx context.Context) ([]*model.Quote, error) {
	return s.repo.GetAllQuotes(ctx)
}

func (s *srv) GetRandomQuote(ctx context.Context) (*model.Quote, error) {
	return s.repo.GetRandomQuote(ctx)
}

func (s *srv) GetQuotesByAuthor(ctx context.Context, quoteAuthor string) ([]*model.Quote, error) {
	var fn = "srv.GetQuotesByAuthor"
	if len(quoteAuthor) == 0 {
		return nil, fmt.Errorf("%s: quote author is empty", fn)
	}

	return s.repo.GetQuotesByAuthor(ctx, quoteAuthor)
}

func (s *srv) DeleteQuote(ctx context.Context, id int64) error {
	var fn = "srv.DeleteQuote"
	if id <= 0 {
		return fmt.Errorf("%s: quote id is invalid", fn)
	}

	return s.repo.DeleteQuote(ctx, id)
}

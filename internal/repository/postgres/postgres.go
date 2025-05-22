package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/solumD/go-quotes-server/internal/model"
	"github.com/solumD/go-quotes-server/internal/repository"
)

type repo struct {
	db *pgxpool.Pool
}

func New(ctx context.Context, path string) (repository.Repository, error) {
	db, err := pgxpool.New(ctx, path)
	if err != nil {
		return nil, err
	}

	return &repo{db: db}, nil
}

func (r *repo) Close() {
	r.db.Close()
}

func (r *repo) SaveQuote(ctx context.Context, quote string, author string) (int64, error) {
	var fn = "repo.SaveQuote"

	q := "INSERT INTO quotes (quote_text, quote_author) VALUES ($1, $2) RETURNING id"

	var id int64
	err := r.db.QueryRow(ctx, q, quote, author).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("%s: failed to save quote: %w", fn, err)
	}

	return id, nil
}

func (r *repo) GetAllQuotes(ctx context.Context) ([]*model.Quote, error) {
	var fn = "repo.GetAllQuotes"

	q := `SELECT id, quote_text, quote_author FROM quotes WHERE is_deleted = false`

	rows, err := r.db.Query(ctx, q)
	if err != nil {
		return nil, fmt.Errorf("%s: failed to get all quotes: %w", fn, err)
	}

	var quotes []*model.Quote
	for rows.Next() {
		var quote model.Quote
		err := rows.Scan(&quote.ID, &quote.Text, &quote.Quthor)
		if err != nil {
			return nil, fmt.Errorf("%s: failed to scan quote: %w", fn, err)
		}
		quotes = append(quotes, &quote)
	}

	return quotes, nil
}

func (r *repo) GetRandomQuote(ctx context.Context) (*model.Quote, error) {
	var fn = "repo.GetRandomQuote"

	q := `SELECT id, quote_text, quote_author FROM quotes WHERE is_deleted = false ORDER BY RANDOM() LIMIT 1`

	var quote model.Quote
	err := r.db.QueryRow(ctx, q).Scan(&quote.ID, &quote.Text, &quote.Quthor)
	if err != nil {
		return nil, fmt.Errorf("%s: failed to get random quote: %w", fn, err)
	}

	return &quote, nil
}

func (r *repo) GetQuotesByAuthor(ctx context.Context, author string) ([]*model.Quote, error) {
	var fn = "repo.GetQuotesByAuthor"

	q := `SELECT id, quote_text, quote_author FROM quotes WHERE is_deleted = false AND quote_author = $1`

	rows, err := r.db.Query(ctx, q, author)
	if err != nil {
		return nil, fmt.Errorf("%s: failed to get quotes by author: %w", fn, err)
	}

	var quotes []*model.Quote
	for rows.Next() {
		var quote model.Quote
		err := rows.Scan(&quote.ID, &quote.Text, &quote.Quthor)
		if err != nil {
			return nil, fmt.Errorf("%s: failed to scan quote: %w", fn, err)
		}
		quotes = append(quotes, &quote)
	}

	return quotes, nil
}

func (r *repo) DeleteQuote(ctx context.Context, id int64) error {
	var fn = "repo.DeleteQuote"

	q := "UPDATE quotes SET is_deleted = true WHERE id = $1"

	_, err := r.db.Exec(ctx, q, id)
	if err != nil {
		return fmt.Errorf("%s: failed to delete quote: %w", fn, err)
	}

	return nil
}

package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/solumD/go-quotes-server/internal/model"
	"github.com/solumD/go-quotes-server/internal/repository"
	repoerrors "github.com/solumD/go-quotes-server/internal/repository/repo_errors"
)

// repo is a struct that contains the database connection.
type repo struct {
	db *pgxpool.Pool
}

// New returns a new repository.
func New(ctx context.Context, path string) (repository.Repository, error) {
	db, err := pgxpool.New(ctx, path)
	if err != nil {
		return nil, err
	}

	return &repo{
		db: db,
	}, nil
}

// Close closes the database connection.
func (r *repo) Close() {
	r.db.Close()
}

// SaveQuote saves a quote in the database.
func (r *repo) SaveQuote(ctx context.Context, quoteText string, quoteAuthor string) (int64, error) {
	var fn = "repo.SaveQuote"

	q := "INSERT INTO quotes (quote_text, quote_author) VALUES ($1, $2) RETURNING id"

	var id int64
	err := r.db.QueryRow(ctx, q, quoteText, quoteAuthor).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("%s: failed to save quote: %w", fn, err)
	}

	return id, nil
}

// GetAllQuotes returns all quotes from the database.
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

// GetRandomQuote returns a random quote from the database.
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

// GetQuotesByAuthor returns quotes by author from the database.
func (r *repo) GetQuotesByAuthor(ctx context.Context, quoteAuthor string) ([]*model.Quote, error) {
	var fn = "repo.GetQuotesByAuthor"

	exist, err := r.isAuthorExists(ctx, quoteAuthor)
	if err != nil {
		return nil, err
	}

	if !exist {
		return nil, repoerrors.ErrAuthorNotExist
	}

	q := `SELECT id, quote_text, quote_author FROM quotes WHERE is_deleted = false AND quote_author = $1`

	rows, err := r.db.Query(ctx, q, quoteAuthor)
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

// DeleteQuote deletes a quote from the database.
func (r *repo) DeleteQuote(ctx context.Context, quoteID int64) error {
	var fn = "repo.DeleteQuote"

	q := "UPDATE quotes SET is_deleted = true WHERE id = $1"

	_, err := r.db.Exec(ctx, q, quoteID)
	if err != nil {
		return fmt.Errorf("%s: failed to delete quote: %w", fn, err)
	}

	return nil
}

// isAuthorExists checks if a quote author exists in the database.
func (r *repo) isAuthorExists(ctx context.Context, quoteAuthor string) (bool, error) {
	var fn = "repo.IsAuthorExists"

	q := "SELECT id FROM quotes WHERE quote_author = $1 AND is_deleted = false"

	var id int64
	err := r.db.QueryRow(ctx, q, quoteAuthor).Scan(&id)
	if err != nil {
		if err.Error() == "no rows in result set" {
			return false, nil
		}

		return false, fmt.Errorf("%s: failed to check if author exists: %w", fn, err)
	}

	return true, nil
}

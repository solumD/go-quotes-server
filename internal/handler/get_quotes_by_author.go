package handler

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/solumD/go-quotes-server/internal/lib/sl"
	"github.com/solumD/go-quotes-server/internal/model"
	repoerrors "github.com/solumD/go-quotes-server/internal/repository/repo_errors"
	srverrors "github.com/solumD/go-quotes-server/internal/service/srv_errors"
)

type getQuotesByAuthorResponse struct {
	Quotes   []*model.Quote `json:"quotes"`
	ErrorMsg string         `json:"error_msg,omitempty"`
}

// GetQuotesByAuthor returns quotes by author.
func (h *handler) GetQuotesByAuthor(ctx context.Context, logger *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var fn = "handler.GetQuotesByAuthor"

		logger = logger.With(
			slog.String("fn", fn),
		)

		author := r.URL.Query().Get("author")
		quotes, err := h.service.GetQuotesByAuthor(ctx, author)
		if err != nil {
			logger.Error("failed to get quotes by author", sl.Err(err))

			w.WriteHeader(http.StatusInternalServerError)

			var resp getQuotesByAuthorResponse
			if err == repoerrors.ErrAuthorNotExist || err == srverrors.ErrAuthorIsEmpty {
				resp.ErrorMsg = err.Error()
			} else {
				resp.ErrorMsg = "failed to get quotes by author"
			}

			data, err := json.Marshal(resp)
			if err != nil {
				logger.Error("failed to marshal response", sl.Err(err))
				return
			}

			w.Write(data)
			return
		}

		w.WriteHeader(http.StatusOK)
		data, err := json.Marshal(getQuotesByAuthorResponse{Quotes: quotes})
		if err != nil {
			logger.Error("failed to marshal response", sl.Err(err))
			return
		}

		w.Write(data)
	}
}

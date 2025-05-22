package handler

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/solumD/go-quotes-server/internal/lib/sl"
	"github.com/solumD/go-quotes-server/internal/model"
)

type GetQuotesByAuthorResponse struct {
	Quotes   []*model.Quote `json:"quotes"`
	ErrorMsg string         `json:"error_msg,omitempty"`
}

func (h *handler) GetQuotesByAuthor(ctx context.Context, logger *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var fn = "handler.GetQuotesByAuthor"

		logger = logger.With(
			slog.String("fn", fn),
			slog.String("request_id", r.Header.Get("X-Request-ID")),
		)

		author := r.URL.Query().Get("author")
		quotes, err := h.service.GetQuotesByAuthor(ctx, author)
		if err != nil {
			logger.Error("failed to get quotes by author", sl.Err(err))

			w.WriteHeader(http.StatusInternalServerError)
			data, err := json.Marshal(GetQuotesByAuthorResponse{ErrorMsg: "failed to get quotes by author"})
			if err != nil {
				logger.Error("failed to marshal response", sl.Err(err))
				return
			}

			w.Write(data)
			return
		}

		w.WriteHeader(http.StatusOK)
		data, err := json.Marshal(GetQuotesByAuthorResponse{Quotes: quotes})
		if err != nil {
			logger.Error("failed to marshal response", sl.Err(err))
			return
		}

		w.Write(data)
	}
}

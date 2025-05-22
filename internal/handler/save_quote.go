package handler

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/solumD/go-quotes-server/internal/lib/sl"
)

type SaveQuoteRequest struct {
	QuoteAuthor string `json:"author"`
	QuoteText   string `json:"quote"`
}

type SaveQuoteResponse struct {
	ID       int64  `json:"id,omitempty"`
	ErrorMsg string `json:"error,omitempty"`
}

func (h *handler) SaveQuote(ctx context.Context, logger *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var fn = "handler.SaveQuote"

		logger = logger.With(
			slog.String("fn", fn),
		)

		var req SaveQuoteRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			logger.Error("failed to decode request", sl.Err(err))

			w.WriteHeader(http.StatusBadRequest)
			data, err := json.Marshal(SaveQuoteResponse{ErrorMsg: "failed to decode request"})
			if err != nil {
				logger.Error("failed to marshal response", sl.Err(err))
				return
			}

			w.Write(data)
			return
		}

		id, err := h.service.SaveQuote(ctx, req.QuoteText, req.QuoteAuthor)
		if err != nil {
			logger.Error("failed to save quote", sl.Err(err))

			w.WriteHeader(http.StatusInternalServerError)
			data, err := json.Marshal(SaveQuoteResponse{ErrorMsg: "failed to save quote"})
			if err != nil {
				logger.Error("failed to marshal response", sl.Err(err))
				return
			}

			w.Write(data)
			return
		}

		w.WriteHeader(http.StatusOK)
		data, err := json.Marshal(SaveQuoteResponse{ID: id})
		if err != nil {
			logger.Error("failed to marshal response", sl.Err(err))
			return
		}

		w.Write(data)
	}
}

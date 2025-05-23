package handler

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/solumD/go-quotes-server/internal/lib/sl"
	srverrors "github.com/solumD/go-quotes-server/internal/service/srv_errors"
)

type saveQuoteRequest struct {
	QuoteAuthor string `json:"author"`
	QuoteText   string `json:"quote"`
}

type saveQuoteResponse struct {
	ID       int64  `json:"id,omitempty"`
	ErrorMsg string `json:"error,omitempty"`
}

// SaveQuote saves a quote.
func (h *handler) SaveQuote(ctx context.Context, logger *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var fn = "handler.SaveQuote"

		logger = logger.With(
			slog.String("fn", fn),
		)

		var req saveQuoteRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			logger.Error("failed to decode request", sl.Err(err))

			w.WriteHeader(http.StatusBadRequest)
			data, err := json.Marshal(saveQuoteResponse{ErrorMsg: "failed to decode request"})
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

			var resp saveQuoteResponse
			if err == srverrors.ErrTextIsEmpty || err == srverrors.ErrAuthorIsEmpty || err == srverrors.ErrTextAndAuthorEmpty {
				resp.ErrorMsg = err.Error()
			} else {
				resp.ErrorMsg = "failed to save quote"
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
		data, err := json.Marshal(saveQuoteResponse{ID: id})
		if err != nil {
			logger.Error("failed to marshal response", sl.Err(err))
			return
		}

		w.Write(data)
	}
}

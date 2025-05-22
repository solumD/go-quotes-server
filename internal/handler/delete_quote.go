package handler

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/solumD/go-quotes-server/internal/lib/sl"
)

type DeleteQuoteRequest struct {
	ID int64 `json:"id"`
}

type DeleteQuoteResponse struct {
	ErrorMsg string `json:"error,omitempty"`
}

func (h *handler) DeleteQuote(ctx context.Context, logger *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var fn = "handler.DeleteQuote"

		logger = logger.With(
			slog.String("fn", fn),
			slog.String("request_id", r.Header.Get("X-Request-ID")),
		)

		var req DeleteQuoteRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			logger.Error("failed to decode request", sl.Err(err))

			w.WriteHeader(http.StatusBadRequest)
			data, err := json.Marshal(DeleteQuoteResponse{ErrorMsg: "failed to decode request"})
			if err != nil {
				logger.Error("failed to marshal response", sl.Err(err))
				return
			}

			w.Write(data)
			return
		}

		err := h.service.DeleteQuote(ctx, req.ID)
		if err != nil {
			logger.Error("failed to delete quote", sl.Err(err))

			w.WriteHeader(http.StatusInternalServerError)
			data, err := json.Marshal(DeleteQuoteResponse{ErrorMsg: "failed to delete quote"})
			if err != nil {
				logger.Error("failed to marshal response", sl.Err(err))
				return
			}

			w.Write(data)
			return
		}

		w.WriteHeader(http.StatusOK)
		data, err := json.Marshal(DeleteQuoteResponse{})
		if err != nil {
			logger.Error("failed to marshal response", sl.Err(err))
			return
		}

		w.Write(data)
	}
}

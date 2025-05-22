package handler

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/solumD/go-quotes-server/internal/lib/sl"
	"github.com/solumD/go-quotes-server/internal/model"
)

type GetRandomQuoteResponse struct {
	Quote    *model.Quote `json:"quote"`
	ErrorMsg string       `json:"error_msg,omitempty"`
}

func (h *handler) GetRandomQuote(ctx context.Context, logger *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var fn = "handler.GetRandomQuote"

		logger = logger.With(
			slog.String("fn", fn),
		)

		quote, err := h.service.GetRandomQuote(ctx)
		if err != nil {
			logger.Error("failed to get random quote", sl.Err(err))

			w.WriteHeader(http.StatusInternalServerError)
			data, err := json.Marshal(GetRandomQuoteResponse{ErrorMsg: "failed to get random quote"})
			if err != nil {
				logger.Error("failed to marshal response", sl.Err(err))
				return
			}

			w.Write(data)
			return
		}

		w.WriteHeader(http.StatusOK)
		data, err := json.Marshal(GetRandomQuoteResponse{Quote: quote})
		if err != nil {
			logger.Error("failed to marshal response", sl.Err(err))
			return
		}

		w.Write(data)
	}
}

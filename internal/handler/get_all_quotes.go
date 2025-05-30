package handler

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/solumD/go-quotes-server/internal/lib/sl"
	"github.com/solumD/go-quotes-server/internal/model"
)

// GetAllQuotesResponse is a struct of the response.
type GetAllQuotesResponse struct {
	Quotes   []*model.Quote `json:"quotes"`
	ErrorMsg string         `json:"error_msg,omitempty"`
}

// GetAllQuotes returns all quotes.
func (h *handler) GetAllQuotes(ctx context.Context, logger *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var fn = "handler.GetAllQuotes"

		logger = logger.With(
			slog.String("fn", fn),
		)

		quotes, err := h.service.GetAllQuotes(ctx)
		if err != nil {
			logger.Error("failed to get all quotes", sl.Err(err))

			w.WriteHeader(http.StatusInternalServerError)
			data, err := json.Marshal(GetAllQuotesResponse{ErrorMsg: "failed to get all quotes"})
			if err != nil {
				logger.Error(ErrMarshalResponse.Error(), sl.Err(err))
				return
			}

			w.Write(data)
			return
		}

		w.WriteHeader(http.StatusOK)
		data, err := json.Marshal(GetAllQuotesResponse{Quotes: quotes})
		if err != nil {
			logger.Error(ErrMarshalResponse.Error(), sl.Err(err))
			return
		}

		w.Write(data)
	}
}

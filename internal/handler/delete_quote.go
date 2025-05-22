package handler

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/solumD/go-quotes-server/internal/lib/sl"
)

type deleteQuoteResponse struct {
	ErrorMsg string `json:"error,omitempty"`
}

// DeleteQuote deletes a quote by ID.
func (h *handler) DeleteQuote(ctx context.Context, logger *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var fn = "handler.DeleteQuote"

		logger = logger.With(
			slog.String("fn", fn),
		)

		id, err := strconv.Atoi(mux.Vars(r)["id"])
		if err != nil {
			logger.Error("failed to get quote id", sl.Err(err))
			w.WriteHeader(http.StatusBadRequest)
			data, err := json.Marshal(deleteQuoteResponse{ErrorMsg: "failed to delete quote"})
			if err != nil {
				logger.Error("failed to marshal response", sl.Err(err))
				return
			}

			w.Write(data)
			return
		}

		err = h.service.DeleteQuote(ctx, int64(id))
		if err != nil {
			logger.Error("failed to delete quote", sl.Err(err))

			w.WriteHeader(http.StatusInternalServerError)
			data, err := json.Marshal(deleteQuoteResponse{ErrorMsg: "failed to delete quote"})
			if err != nil {
				logger.Error("failed to marshal response", sl.Err(err))
				return
			}

			w.Write(data)
			return
		}

		w.WriteHeader(http.StatusOK)
		data, err := json.Marshal(deleteQuoteResponse{})
		if err != nil {
			logger.Error("failed to marshal response", sl.Err(err))
			return
		}

		w.Write(data)
	}
}

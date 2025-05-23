package handler

import (
	"context"
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/solumD/go-quotes-server/internal/lib/sl"
	srverrors "github.com/solumD/go-quotes-server/internal/service/srv_errors"
)

var (
	// ErrFailedToGetQuoteID is an error that is returned when the quote id is not found.
	ErrFailedToGetQuoteID = errors.New("failed to get quote id")
)

// DeleteQuoteResponse is a struct of the response.
type DeleteQuoteResponse struct {
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
			data, err := json.Marshal(DeleteQuoteResponse{ErrorMsg: ErrFailedToGetQuoteID.Error()})
			if err != nil {
				logger.Error(ErrMarshalResponse.Error(), sl.Err(err))
				return
			}

			w.Write(data)
			return
		}

		err = h.service.DeleteQuote(ctx, int64(id))
		if err != nil {
			logger.Error("failed to delete quote", sl.Err(err))

			var resp DeleteQuoteResponse
			if err == srverrors.ErrInvalidID {
				w.WriteHeader(http.StatusBadRequest)
				resp.ErrorMsg = err.Error()
			} else {
				w.WriteHeader(http.StatusInternalServerError)
				resp.ErrorMsg = "failed to delete quote"
			}

			data, err := json.Marshal(resp)
			if err != nil {
				logger.Error(ErrMarshalResponse.Error(), sl.Err(err))
				return
			}

			w.Write(data)
			return
		}

		w.WriteHeader(http.StatusOK)
		data, err := json.Marshal(DeleteQuoteResponse{})
		if err != nil {
			logger.Error(ErrMarshalResponse.Error(), sl.Err(err))
			return
		}

		w.Write(data)
	}
}

package handler

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/solumD/go-quotes-server/internal/lib/sl"
	srverrors "github.com/solumD/go-quotes-server/internal/service/srv_errors"
)

// SaveQuoteRequest is a struct of the request.
type SaveQuoteRequest struct {
	QuoteAuthor string `json:"author"`
	QuoteText   string `json:"quote"`
}

// SaveQuoteResponse is a struct of the response.
type SaveQuoteResponse struct {
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

		var req SaveQuoteRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			logger.Error(ErrUnmarshalRequest.Error(), sl.Err(err))

			w.WriteHeader(http.StatusBadRequest)
			data, err := json.Marshal(SaveQuoteResponse{ErrorMsg: ErrUnmarshalRequest.Error()})
			if err != nil {
				logger.Error(ErrMarshalResponse.Error(), sl.Err(err))
				return
			}

			w.Write(data)
			return
		}

		id, err := h.service.SaveQuote(ctx, req.QuoteText, req.QuoteAuthor)
		if err != nil {
			logger.Error("failed to save quote", sl.Err(err))

			var resp SaveQuoteResponse
			if err == srverrors.ErrTextIsEmpty || err == srverrors.ErrAuthorIsEmpty || err == srverrors.ErrTextAndAuthorEmpty {
				w.WriteHeader(http.StatusBadRequest)
				resp.ErrorMsg = err.Error()
			} else {
				w.WriteHeader(http.StatusInternalServerError)
				resp.ErrorMsg = "failed to save quote"
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
		data, err := json.Marshal(SaveQuoteResponse{ID: id})
		if err != nil {
			logger.Error(ErrMarshalResponse.Error(), sl.Err(err))
			return
		}

		w.Write(data)
	}
}

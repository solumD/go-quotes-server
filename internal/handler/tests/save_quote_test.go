package tests

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/solumD/go-quotes-server/internal/handler"
	"github.com/solumD/go-quotes-server/internal/lib/sl"
	mock_service "github.com/solumD/go-quotes-server/internal/service/mocks"
	srverrors "github.com/solumD/go-quotes-server/internal/service/srv_errors"
)

func TestSaveQuote(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockService := mock_service.NewMockService(mockCtrl)
	testCases := []struct {
		name               string
		req                handler.SaveQuoteRequest
		mockError          error
		expectedError      error
		expectedStatusCode int
	}{
		{
			name: "success",
			req: handler.SaveQuoteRequest{
				QuoteAuthor: "author",
				QuoteText:   "quote",
			},
			mockError:          nil,
			expectedError:      nil,
			expectedStatusCode: http.StatusOK,
		},
		{
			name: "error quote text is empty",
			req: handler.SaveQuoteRequest{
				QuoteAuthor: "author",
				QuoteText:   "",
			},
			mockError:          srverrors.ErrTextIsEmpty,
			expectedError:      srverrors.ErrTextIsEmpty,
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name: "error quote author is empty",
			req: handler.SaveQuoteRequest{
				QuoteAuthor: "",
				QuoteText:   "quote",
			},
			mockError:          srverrors.ErrAuthorIsEmpty,
			expectedError:      srverrors.ErrAuthorIsEmpty,
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name: "error both quote text and author are empty",
			req: handler.SaveQuoteRequest{
				QuoteAuthor: "",
				QuoteText:   "",
			},
			mockError:          srverrors.ErrTextAndAuthorEmpty,
			expectedError:      srverrors.ErrTextAndAuthorEmpty,
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name: "error saving quote",
			req: handler.SaveQuoteRequest{
				QuoteAuthor: "author",
				QuoteText:   "quote",
			},
			mockError:          errors.New("srv error"),
			expectedError:      errors.New("failed to save quote"),
			expectedStatusCode: http.StatusInternalServerError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockService.EXPECT().SaveQuote(gomock.Any(), tc.req.QuoteText, tc.req.QuoteAuthor).Return(int64(0), tc.mockError)

			h := handler.New(mockService).SaveQuote(context.Background(), sl.NewDiscardLogger())

			reqBody := fmt.Sprintf(`{"quote": "%s", "author": "%s"}`, tc.req.QuoteText, tc.req.QuoteAuthor)
			req, _ := http.NewRequest(http.MethodPost, "/quotes", bytes.NewReader([]byte(reqBody)))

			recorder := httptest.NewRecorder()
			h.ServeHTTP(recorder, req)

			if tc.expectedError == nil {
				if tc.expectedStatusCode != recorder.Code {
					t.Errorf("expected status code: %d | but got: %d", tc.expectedStatusCode, recorder.Code)
				}
			} else {
				if tc.expectedStatusCode != recorder.Code {
					t.Errorf("expected status code: %d | but got: %d", tc.expectedStatusCode, recorder.Code)
				}

				var resp handler.SaveQuoteResponse
				json.Unmarshal(recorder.Body.Bytes(), &resp)
				if tc.expectedError.Error() != resp.ErrorMsg {
					t.Errorf("expected error message: %s | but got: %s", tc.expectedError.Error(), resp.ErrorMsg)
				}
			}
		})
	}
}

package tests

import (
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
	"github.com/solumD/go-quotes-server/internal/model"
	mock_service "github.com/solumD/go-quotes-server/internal/service/mocks"
)

func TestGetRandomQuote(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockService := mock_service.NewMockService(mockCtrl)

	testCases := []struct {
		name               string
		mockError          error
		expectedError      error
		expectedStatusCode int
	}{
		{
			name:               "success",
			mockError:          nil,
			expectedError:      nil,
			expectedStatusCode: http.StatusOK,
		},
		{
			name:               "error failed to get random quote",
			mockError:          errors.New("srv error"),
			expectedError:      errors.New("failed to get random quote"),
			expectedStatusCode: http.StatusInternalServerError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockService.EXPECT().GetRandomQuote(gomock.Any()).Return(&model.Quote{}, tc.mockError)

			h := handler.New(mockService).GetRandomQuote(context.Background(), sl.NewDiscardLogger())

			req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/quotes/random"), nil)

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

				var resp handler.GetQuotesByAuthorResponse
				json.Unmarshal(recorder.Body.Bytes(), &resp)
				if tc.expectedError.Error() != resp.ErrorMsg {
					t.Errorf("expected error message: %s | but got: %s", tc.expectedError.Error(), resp.ErrorMsg)
				}
			}
		})
	}
}

package tests

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/solumD/go-quotes-server/internal/handler"
	"github.com/solumD/go-quotes-server/internal/lib/sl"
	mock_service "github.com/solumD/go-quotes-server/internal/service/mocks"
	srverrors "github.com/solumD/go-quotes-server/internal/service/srv_errors"
)

func TestDeleteQuote(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockService := mock_service.NewMockService(mockCtrl)

	testCases := []struct {
		name               string
		id                 int64
		mockError          error
		expectedError      error
		expectedStatusCode int
	}{
		{
			name:               "success",
			id:                 1,
			mockError:          nil,
			expectedError:      nil,
			expectedStatusCode: http.StatusOK,
		},
		{
			name:               "error quote id is invalid",
			id:                 -1,
			mockError:          srverrors.ErrInvalidID,
			expectedError:      srverrors.ErrInvalidID,
			expectedStatusCode: http.StatusBadRequest,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockService.EXPECT().DeleteQuote(gomock.Any(), tc.id).Return(tc.mockError)

			h := handler.New(mockService).DeleteQuote(context.Background(), sl.NewDiscardLogger())

			vars := map[string]string{
				"id": strconv.FormatInt(tc.id, 10),
			}
			req, _ := http.NewRequest(http.MethodDelete, "/quotes/{id}", nil)
			req = mux.SetURLVars(req, vars)

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

				var resp handler.DeleteQuoteResponse
				json.Unmarshal(recorder.Body.Bytes(), &resp)
				if tc.expectedError.Error() != resp.ErrorMsg {
					t.Errorf("expected error message: %s | but got: %s", tc.expectedError.Error(), resp.ErrorMsg)
				}
			}
		})
	}
}

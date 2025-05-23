package tests

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/solumD/go-quotes-server/internal/handler"
	"github.com/solumD/go-quotes-server/internal/lib/sl"
	"github.com/solumD/go-quotes-server/internal/model"
	repoerrors "github.com/solumD/go-quotes-server/internal/repository/repo_errors"
	mock_service "github.com/solumD/go-quotes-server/internal/service/mocks"
	srverrors "github.com/solumD/go-quotes-server/internal/service/srv_errors"
)

func TestGetQuotesByAuthor(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockService := mock_service.NewMockService(mockCtrl)

	testCases := []struct {
		name               string
		author             string
		mockError          error
		expectedError      error
		expectedStatusCode int
	}{
		{
			name:               "success",
			author:             "some_author",
			mockError:          nil,
			expectedError:      nil,
			expectedStatusCode: http.StatusOK,
		},
		{
			name:               "error author is empty",
			author:             "",
			mockError:          srverrors.ErrAuthorIsEmpty,
			expectedError:      srverrors.ErrAuthorIsEmpty,
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name:               "error author not exists",
			author:             "some_author",
			mockError:          repoerrors.ErrAuthorNotExist,
			expectedError:      repoerrors.ErrAuthorNotExist,
			expectedStatusCode: http.StatusBadRequest,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockService.EXPECT().GetQuotesByAuthor(gomock.Any(), tc.author).Return([]*model.Quote{}, tc.mockError)

			h := handler.New(mockService).GetQuotesByAuthor(context.Background(), sl.NewDiscardLogger())

			req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/quotes?author=%s", tc.author), nil)

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

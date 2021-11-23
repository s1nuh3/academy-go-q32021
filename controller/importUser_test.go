package controller

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/s1nuh3/academy-go-q32021/model"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

//ImportHandlerMock - struc to mock the handler
type ImportHandlerMock struct {
	mock.Mock
}

//ImportUserUC - stub
func (cm *ImportHandlerMock) ImportUserUC(id int) (*model.Users, error) {
	args := cm.Called(id)
	return args.Get(0).(*model.Users), args.Error(1)
}

//TestImportHandler - Apply unit test over the controller/handler of impor user from external API
func TestImportHandler(t *testing.T) {

	testCases := []struct {
		name        string
		httpStatus  int
		hasError    bool
		body        string
		ID          string
		expectedUC  *model.Users
		errorUC     error
		testInvalid bool
	}{
		{
			name:        "Invalid user ID",
			httpStatus:  http.StatusBadRequest,
			hasError:    true,
			body:        "ID provided is not valid",
			ID:          "0",
			expectedUC:  nil,
			errorUC:     nil,
			testInvalid: true,
		},
		{
			name:        "Import successful",
			httpStatus:  http.StatusOK,
			hasError:    false,
			body:        "User",
			ID:          "989",
			expectedUC:  &model.Users{ID: 989, Name: "Devi Malik", Email: "malik_devi@rowe.com", Gender: "female", Status: "inactive"},
			errorUC:     nil,
			testInvalid: false,
		},
		{
			name:        "Import fails at getting user",
			httpStatus:  http.StatusNotFound,
			hasError:    true,
			body:        "Any error",
			ID:          "989",
			expectedUC:  &model.Users{ID: 989, Name: "Devi Malik", Email: "malik_devi@rowe.com", Gender: "female", Status: "inactive"},
			errorUC:     errors.New("Any error"),
			testInvalid: false,
		},
		{
			name:        "Import fails at getting user user ID 0",
			httpStatus:  http.StatusNotFound,
			hasError:    true,
			body:        "ID not found in external API",
			ID:          "989",
			expectedUC:  &model.Users{ID: 0, Name: "Devi Malik", Email: "malik_devi@rowe.com", Gender: "female", Status: "inactive"},
			errorUC:     nil,
			testInvalid: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mock := ImportHandlerMock{}
			intID, err := strconv.Atoi(tc.ID)
			assert.NoError(t, err)
			mock.On("ImportUserUC", intID).Return(tc.expectedUC, tc.errorUC)
			uc := NewImportHandler(&mock)

			if tc.testInvalid {
				tc.ID = "test"
			}

			req, err := http.NewRequest(http.MethodGet, "/user/import/"+tc.ID, nil)
			req.Close = true
			assert.NoError(t, err)

			rr := httptest.NewRecorder()
			r := mux.NewRouter()
			r.HandleFunc("/user/import/{id}", uc.ImportHdl)
			r.ServeHTTP(rr, req)
			assert.Equal(t, tc.httpStatus, rr.Code)
			respBody, err := json.Marshal(tc.body)
			if tc.body == "User" {
				respBody, err = json.Marshal(tc.expectedUC)
			}
			assert.Equal(t, respBody, rr.Body.Bytes())
			assert.Nil(t, err)
		})
	}
}

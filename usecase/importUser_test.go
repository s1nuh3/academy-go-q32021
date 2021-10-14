package usecase

import (
	"errors"
	"testing"

	"github.com/s1nuh3/academy-go-q32021/model"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

//ClientAPIMock - Struc to mock the implmentation of the client API
type ClientAPIMock struct {
	mock.Mock
}

//ImportUser - mock implementation
func (cm *ClientAPIMock) ImportUser(id int) (*model.Users, error) {
	args := cm.Called(id)
	return args.Get(0).(*model.Users), args.Error(1)
}

//RepoServiceMock - Struc to mock the implementation of the repo
type RepoServiceMock struct {
	mock.Mock
}

//Get - mock implementation
func (r *RepoServiceMock) Get(id int) (*model.Users, error) {
	args := r.Called(id)
	return args.Get(0).(*model.Users), args.Error(1)
}

//List - mock implementation
func (r *RepoServiceMock) List() (*[]model.Users, error) {
	args := r.Called()
	return args.Get(0).(*[]model.Users), args.Error(1)
}

//TestUseCaseImporUser_Test - Unit test for importing a user from client API
func TestUseCaseImportUser(t *testing.T) {
	testCases := []struct {
		name        string
		expectedCli *model.Users
		expectedRp  *model.Users
		hasErrorCli bool
		errorCli    error
		hasErrorRp  bool
		errorRp     error
		ID          int
	}{
		{
			name:        "Successful Import",
			expectedCli: &model.Users{ID: 45, Name: "Test"},
			expectedRp:  &model.Users{ID: 0, Name: ""},
			hasErrorCli: false,
			errorCli:    nil,
			hasErrorRp:  false,
			errorRp:     nil,
			ID:          45,
		},
		{
			name:        "User already exist",
			expectedCli: nil,
			expectedRp:  &model.Users{ID: 46, Name: "Test"},
			hasErrorCli: true,
			errorCli:    errors.New("el ID de usuario ya existe"),
			hasErrorRp:  false,
			errorRp:     nil,
			ID:          46,
		},
		{
			name:        "Error on get repo",
			expectedCli: nil,
			expectedRp:  nil,
			hasErrorCli: false,
			errorCli:    nil,
			hasErrorRp:  true,
			errorRp:     errors.New("Not nil"),
			ID:          46,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockClient := ClientAPIMock{}
			mockRepo := RepoServiceMock{}
			mockRepo.On("Get", tc.ID).Return(tc.expectedRp, tc.errorRp)
			mockClient.On("ImportUser", tc.ID).Return(tc.expectedCli, tc.errorCli)

			uc := NewImportUser(&mockClient, &mockRepo)

			imported, err := uc.ImportUserUC(tc.ID)
			assert.EqualValues(t, tc.expectedCli, imported)

			if tc.hasErrorCli {
				assert.EqualError(t, err, tc.errorCli.Error())
			} else if tc.hasErrorRp {
				assert.Error(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

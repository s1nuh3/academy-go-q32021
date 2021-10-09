package usecase

import (
	"errors"
	"testing"

	"github.com/s1nuh3/academy-go-q32021/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type ImportHandlerMock struct {
	mock.Mock
}

func (cm *ImportHandlerMock) ImportUser(id int) (*model.Users, error) {
	args := cm.Called(id)
	return args.Get(0).(*model.Users), args.Error(1)
}

type RepoServiceMock struct {
	mock.Mock
}

func (r *RepoServiceMock) Get(id int) (*model.Users, error) {
	args := r.Called(id)
	return args.Get(0).(*model.Users), args.Error(1)
}

//Get the list of users
func (r *RepoServiceMock) List() (*[]model.Users, error) {
	args := r.Called()
	return args.Get(0).(*[]model.Users), args.Error(1)
}

func TestUseCaseImporUser_Test(t *testing.T) {
	testCases := []struct {
		name        string
		expectedHdl *model.Users
		expectedRp  *model.Users
		hasError    bool
		error       error
		ID          int
	}{
		{
			name:        "Successful Import",
			expectedHdl: &model.Users{ID: 45, Name: "Test"},
			expectedRp:  &model.Users{ID: 0, Name: ""},
			hasError:    false,
			error:       nil,
			ID:          45,
		},
		{
			name:        "Failed Import",
			expectedHdl: nil,
			expectedRp:  &model.Users{ID: 46, Name: "Test"},
			hasError:    true,
			error:       errors.New("el ID de usuario ya existe"),
			ID:          46,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockHandler := ImportHandlerMock{}
			mockRepo := RepoServiceMock{}
			mockRepo.On("Get", tc.ID).Return(tc.expectedRp, tc.error)
			mockHandler.On("ImportUser", tc.ID).Return(tc.expectedHdl, tc.error)

			uc := NewConsumeAPI(&mockHandler, &mockRepo)

			imported, err := uc.ImportUserCtrl(tc.ID)
			assert.EqualValues(t, tc.expectedHdl, imported)

			if tc.hasError {
				assert.EqualError(t, err, tc.error.Error())
			} else {
				assert.Nil(t, err)
			}
		})
	}
	// testCasespUserService.On("ImportUserCtrl", 45).Return(&model.Users{ID: 1, Name: "Test"}, nil)
	// contro := NewImportCtrl(&ImpUserService)
	// imported, err := contro.us.ImportUserCtrl(45)
	// assertions.NoError(err, "No user was imported ")
	// assertions.Equal(&model.Users{ID: 1, Name: "Test"}, imported, "Incorrect user")
}

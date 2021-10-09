package usecase

import (
	"errors"
	"log"

	"github.com/s1nuh3/academy-go-q32021/model"
)

//CtrlImpUser - To implement the ConsumeAPI interface
type UseCaseUser struct {
	client ConsumeAPI
	repo   Repository
}

//NewImport create service pass it to Controller
func NewConsumeAPI(c ConsumeAPI, r Repository) *UseCaseUser {
	return &UseCaseUser{
		client: c,
		repo:   r,
	}
}

// GetExternalUserCtrl - Returns a user by id if it's found in a csv file
func (s *UseCaseUser) ImportUserCtrl(id int) (*model.Users, error) {

	user, err := s.repo.Get(id)
	if err != nil {
		log.Fatal(err)
	}
	if user.ID != 0 {
		return nil, errors.New("el ID de usuario ya existe")
	}

	user, err = s.client.ImportUser(id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

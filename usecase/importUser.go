package usecase

import (
	"errors"
	"log"

	"github.com/s1nuh3/academy-go-q32021/model"
)

//UseCaseImportUser - To implement the controller contract
type UseCaseImportUser struct {
	client ClientAPI
	repo   Repository
}

//NewImportUser - create a new usecase to be consumed by controller, receives client api and repo
func NewImportUser(c ClientAPI, r Repository) *UseCaseImportUser {
	return &UseCaseImportUser{
		client: c,
		repo:   r,
	}
}

//ImportUserUC - Import a new user into the CSV file, if the ID doesn't already exist
func (s *UseCaseImportUser) ImportUserUC(id int) (*model.Users, error) {

	user, err := s.repo.Get(id)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	if user.ID != 0 {
		return nil, errors.New("el ID de usuario ya existe")
	}

	user, err = s.client.ImportUser(id)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return user, nil
}

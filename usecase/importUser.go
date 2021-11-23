package usecase

import (
	"errors"

	"github.com/s1nuh3/academy-go-q32021/model"
)

//UseCaseImportUser - To implement the controller contract
type UseCaseImportUser struct {
	client ClientAPI
	repo   Repository
}

//NewImportUser - Creates an new instance for to be consumed by controller, receives client api and repo
func NewImportUser(c ClientAPI, r Repository) *UseCaseImportUser {
	return &UseCaseImportUser{client: c, repo: r}
}

//ImportUserUC - Import a new user into the CSV file, if the ID doesn't already exist
func (ui *UseCaseImportUser) ImportUserUC(id int) (*model.Users, error) {
	user, err := ui.repo.Get(id)
	if err != nil {
		return nil, err
	}
	if user.ID != 0 {
		return nil, errors.New("el ID de usuario ya existe")
	}
	user, err = ui.client.ImportUser(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

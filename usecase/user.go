package usecase

import "github.com/s1nuh3/academy-go-q32021/model"

//UseCaseUser - Struc to implement the reposotory interface
type UseCaseUser struct {
	repo Repository
}

//NewService - Creates an new instance to be cosumen at handler
func NewUser(r Repository) *UseCaseUser {
	return &UseCaseUser{repo: r}
}

// ListUsers - Returns a colection of model.users from a csv file
func (s *UseCaseUser) ListUsers() (*[]model.Users, error) {
	return s.repo.List()
}

// GetUser - Returns a user by id if it's found in a csv file
func (s *UseCaseUser) GetUser(id int) (*model.Users, error) {
	return s.repo.Get(id)
}

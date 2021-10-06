package usecase

import "github.com/s1nuh3/academy-go-q32021/model"

type UserUC struct {
	repo Repository
}

//NewService create service for user usecase
func NewUser(r Repository) *UserUC {
	return &UserUC{
		repo: r,
	}
}

// ListUsers - Returns a colection of model.users from a csv file
func (s *UserUC) ListUsers() (*[]model.Users, error) {
	return s.repo.List()
}

// GetUser - Returns a user by id if it's found in a csv file
func (s *UserUC) GetUser(id int) (*model.Users, error) {
	return s.repo.Get(id)
}

// GetUser - Returns a user by id if it's found in a csv file
func (s *UserUC) CreateUser(name, email, gender string, status string) (model.Users, error) {
	return s.repo.Create(name, email, gender, status)
}

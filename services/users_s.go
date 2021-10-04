package services

import (
	"github.com/s1nuh3/academy-go-q32021/models"
)

type UserService struct {
	repo Repository
}

//NewService create service for user usecase
func NewService(r Repository) *UserService {
	return &UserService{
		repo: r,
	}
}

// ListUsers - Returns a colection of model.users from a csv file
func (s *UserService) ListUsers() (*[]models.Users, error) {
	return s.repo.List()
}

// GetUser - Returns a user by id if it's found in a csv file
func (s *UserService) GetUser(id int) (*models.Users, error) {
	return s.repo.Get(id)
}

// GetUser - Returns a user by id if it's found in a csv file
func (s *UserService) CreateUser(name, email, gender string, status bool) (models.Users, error) {
	return s.repo.Create(name, email, gender, status)
}

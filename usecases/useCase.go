package usecases

import "github.com/s1nuh3/academy-go-q32021/models"

//UseCase interface
type UserUseCase interface {
	GetUser(id int) (*models.Users, error)
	ListUsers() (*[]models.Users, error)
	CreateUser(name, email, gender string, status bool) (models.Users, error)
}

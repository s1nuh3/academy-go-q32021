package usecase

import "github.com/s1nuh3/academy-go-q32021/model"

//UseCase interface
type UserUseCase interface {
	GetUser(id int) (*model.Users, error)
	ListUsers() (*[]model.Users, error)
	CreateUser(name, email, gender string, status string) (model.Users, error)
}

type Reader interface {
	Get(id int) (*model.Users, error)
	List() (*[]model.Users, error)
}

//Writer user writer
type Writer interface {
	Create(name, email, gender string, status string) (model.Users, error)
	//Delete(id int) error
}

//Repository interface
type Repository interface {
	Reader
	Writer
}

type ConsumeAPI interface {
	ImportUser(id int) (*model.Users, error)
}

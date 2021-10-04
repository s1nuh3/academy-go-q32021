package services

import "github.com/s1nuh3/academy-go-q32021/models"

//Reader interface
type Reader interface {
	Get(id int) (*models.Users, error)
	List() (*[]models.Users, error)
}

//Writer user writer
type Writer interface {
	Create(name, email, gender string, status bool) (models.Users, error)
	//Delete(id int) error
}

//Repository interface
type Repository interface {
	Reader
	Writer
}

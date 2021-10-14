package usecase

import (
	"github.com/s1nuh3/academy-go-q32021/model"
)

//Reader interface
type Reader interface {
	Get(id int) (*model.Users, error)
	List() (*[]model.Users, error)
}

//Repository interface
type Repository interface {
	Reader
}

//ClientAPI interface
type ClientAPI interface {
	ImportUser(id int) (*model.Users, error)
}

//GoRoutine interface
type GoRoutine interface {
	WorkPool(filter, items, itemsPerWorker, workers int) (*[]model.Users, error)
}

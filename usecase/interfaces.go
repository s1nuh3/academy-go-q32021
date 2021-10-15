package usecase

import (
	"github.com/s1nuh3/academy-go-q32021/model"
)

//Reader Contract
type Reader interface {
	Get(id int) (*model.Users, error)
	List() (*[]model.Users, error)
}

//Repository Contract
type Repository interface {
	Reader
}

//ClientAPI Contract
type ClientAPI interface {
	ImportUser(id int) (*model.Users, error)
}

//GoRoutine Contract
type GoRoutine interface {
	WorkPool(filter, items, itemsPerWorker, workers int) (*[]model.Users, error)
}

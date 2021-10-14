package usecase

import (
	"math"

	"github.com/s1nuh3/academy-go-q32021/model"
)

//UseCaseUser - Struc to implement the reposotory interface
type UseCaseGoRoutine struct {
	gr GoRoutine
}

//NewGoRoutine create service for user usecase
func NewGoRoutine(gr GoRoutine) *UseCaseGoRoutine {
	return &UseCaseGoRoutine{
		gr: gr,
	}
}

// ListUsers - Returns a colection of model.users from a csv file
func (ur *UseCaseGoRoutine) ReadConcurrent(filter, items, itemsPerWorker int) (*[]model.Users, error) {
	workers := int(math.Ceil(float64(items) / float64(itemsPerWorker)))
	return ur.gr.WorkPool(filter, items, itemsPerWorker, workers)
}

package user

import (
	"errors"
	"fmt"
	"log"
	"strconv"

	"github.com/s1nuh3/academy-go-q32021/model"
)

//RepoService - To implment the repository over CSV File
type RepoService struct {
	c CSV
}

//CSV - Contract to access users CSV File
type CSV interface {
	GetData() ([][]string, error)
}

//New - Creates an new instance to a access Users, receive repo csv
func New(c CSV) *RepoService {
	return &RepoService{c: c}
}

//Get - Returns a user by ID if it's found
func (r *RepoService) Get(id int) (*model.Users, error) {
	return getUser(id, r.c)
}

//List - Returns the complete list of users
func (r *RepoService) List() (*[]model.Users, error) {
	return listUsers(r.c)
}

//ParseUserRecord -- Parse a slice of string into a user model strct
func ParseUserRecord(r []string) (*model.Users, error) {
	id, err := strconv.Atoi(r[0])
	if err != nil {
		return nil, errors.New("invalid record")
	}
	//status, _ := strconv.ParseBool(r[4])
	u := model.Users{
		ID:     id,
		Name:   r[1],
		Email:  r[2],
		Gender: r[3],
		Status: r[4],
	}
	return &u, nil
}

func listUsers(c CSV) (*[]model.Users, error) {
	rcd, err := c.GetData()
	if err != nil {
		log.Print(err.Error())
		return nil, err
	}
	result, _, err := convertUsers(rcd)
	return result, err
}

func convertUsers(rcd [][]string) (*[]model.Users, int, error) {
	var invalidRecords int
	var u []model.Users
	for _, r := range rcd {
		data, err := ParseUserRecord(r)
		if err != nil {
			invalidRecords++
		} else {
			u = append(u, *data)
		}
	}
	if invalidRecords > 0 {
		fmt.Println("Files has invalid records - #total: ", invalidRecords)
	}
	return &u, invalidRecords, nil
}

func getUser(id int, c CSV) (*model.Users, error) {
	u := &model.Users{}
	rcd, err := c.GetData()
	if err != nil {
		log.Print(err.Error())
		return u, err
	}

	for _, r := range rcd {
		if r[0] == strconv.Itoa(id) {
			u, _ = ParseUserRecord(r)
			break
		}
	}
	return u, nil
}

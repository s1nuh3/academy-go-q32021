package user

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/s1nuh3/academy-go-q32021/model"
)

//Repository - To implment the repository over CSV File
type Repository struct {
	c   CSV
	err error
}

type CSV interface {
	GetData() ([][]string, error)
	WriteALLData(records [][]string) error
	WriteRowData(record []string) error
	ValidateFile() error
}

//NewRepo - Creates the a new repository implementation
func New(c CSV) *Repository {
	err := c.ValidateFile()

	if err != nil {
		return &Repository{
			c:   c,
			err: err,
		}
	}

	return &Repository{c: c, err: nil}
}

//Create an user
func (r *Repository) Create(name, email, gender string, status string) (model.Users, error) {
	// TODO implment
	rand.Seed(time.Now().UnixNano())
	NewId := rand.Int()
	nu := model.Users{
		ID:     NewId,
		Name:   name,
		Email:  email,
		Gender: gender,
		Status: status,
	}

	return nu, nil
}

//Get an user
func (r *Repository) Get(id int) (*model.Users, error) {
	return getUser(id, r.c)
}

func getUser(id int, c CSV) (*model.Users, error) {
	u := &model.Users{}
	rcd, err := c.GetData()
	if err != nil {
		return u, err
	}

	for _, r := range rcd {
		if r[0] == strconv.Itoa(id) {
			u, _ = parseUserRecord(r)
			break
		}
	}
	return u, nil
}

//Get the list of users
func (r *Repository) List() (*[]model.Users, error) {
	return listUsers(r.c)
}

func listUsers(c CSV) (*[]model.Users, error) {
	rcd, err := c.GetData()
	if err != nil {
		return nil, err
	}
	result, _, err := convertUsers(rcd)
	return result, err
}

func convertUsers(rcd [][]string) (*[]model.Users, int, error) {
	var invalidRecords int
	var u []model.Users
	for _, r := range rcd {
		data, err := parseUserRecord(r)
		if err != nil {
			invalidRecords++
		} else {
			u = append(u, *data)
		}
	}
	fmt.Println("Files has invalid records - #total: ", invalidRecords)
	return &u, invalidRecords, nil
}

func parseUserRecord(r []string) (*model.Users, error) {
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

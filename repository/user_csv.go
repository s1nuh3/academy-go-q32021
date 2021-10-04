package repository

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/s1nuh3/academy-go-q32021/common"
	"github.com/s1nuh3/academy-go-q32021/models"
)

//UserCSV mysql repo
type UserCSV struct {
	db  string
	err error
}

//NewUserCSV create new repository
func NewUserCSV(db string) *UserCSV {
	err := common.ValidateFile(db)

	if err != nil {
		return &UserCSV{
			db:  db,
			err: err,
		}
	}

	return &UserCSV{
		db:  db,
		err: nil,
	}
}

//Create an user
func (r *UserCSV) Create(name, email, gender string, status bool) (models.Users, error) {
	// TODO implment
	rand.Seed(time.Now().UnixNano())
	NewId := rand.Int()
	nu := models.Users{
		ID:     NewId,
		Name:   name,
		Email:  email,
		Gender: gender,
		Status: status,
	}

	return nu, nil
}

//Get an user
func (r *UserCSV) Get(id int) (*models.Users, error) {
	return getUser(id, r.db)
}

func getUser(id int, db string) (*models.Users, error) {
	u := &models.Users{}
	rcd, err := common.GetData(db)
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

//Get an user
func (r *UserCSV) List() (*[]models.Users, error) {
	return listUsers(r.db)
}

//List users
func listUsers(db string) (*[]models.Users, error) {
	rcd, err := common.GetData(db)
	if err != nil {
		return nil, err
	}
	result, _, err := convertUsers(rcd)
	return result, err
}

func convertUsers(rcd [][]string) (*[]models.Users, int, error) {
	var invalidRecords int
	var u []models.Users
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

func parseUserRecord(r []string) (*models.Users, error) {
	id, err := strconv.Atoi(r[0])
	if err != nil {
		return nil, errors.New("invalid record")
	}
	status, _ := strconv.ParseBool(r[4])
	u := models.Users{
		ID:     id,
		Name:   r[1],
		Email:  r[2],
		Gender: r[3],
		Status: status,
	}
	return &u, nil
}

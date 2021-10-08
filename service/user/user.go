package user

import (
	"errors"
	"fmt"
	"log"
	"strconv"

	"github.com/s1nuh3/academy-go-q32021/model"
)

//RepoImpCSV - To implment the repository over CSV File
type RepoImpCSV struct {
	c CSV
}

type CSV interface {
	GetData() ([][]string, error)
	WriteALLData(records [][]string) error
	WriteRowData(record []string) error
}

//NewRepo - Creates the a new repository implementation
func New(c CSV) *RepoImpCSV {
	return &RepoImpCSV{c: c}
}

//Get an user
func (r *RepoImpCSV) Get(id int) (*model.Users, error) {
	return getUser(id, r.c)
}

//Get the list of users
func (r *RepoImpCSV) List() (*[]model.Users, error) {
	return listUsers(r.c)
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
		data, err := parseUserRecord(r)
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
			u, _ = parseUserRecord(r)
			break
		}
	}
	return u, nil
}

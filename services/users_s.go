package services

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/s1nuh3/academy-go-q32021/common"
	"github.com/s1nuh3/academy-go-q32021/models"
)

// GetUsersfromCSV - Returns a colection of model.users from a csv file
func GetUsersfromCSV() (*[]models.Users, error) {
	var invalidRecords int
	rcd, err := common.ReadCsv("./repositories/files/usersdata.csv")
	if err != nil {
		return nil, err
	}
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
	return &u, nil
}

// GetUserbyIdfromCSV - Returns a user by id if it's found in a csv file
func GetUserbyIdfromCSV(id int) (*models.Users, error) {
	u := &models.Users{}
	rcd, err := common.ReadCsv("./repositories/files/usersdata.csv")
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

package clientAPI

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strconv"

	"github.com/s1nuh3/academy-go-q32021/config"
	"github.com/s1nuh3/academy-go-q32021/model"

	resty "github.com/go-resty/resty/v2"
)

//ClientService - Struc to implement the Client API
type ClientService struct {
	client resty.Client
	csv    CSV
}

//CSV - Contract to write the imported users to CSV File
type CSV interface {
	WriteALLData(records [][]string) error
	WriteRowData(record []string) error
}

// Model to deposit the response form the client API
type extUser struct {
	Meta interface{} `json:"meta"`
	Data model.Users `json:"data"`
}

//New - Creates an instance to ConsumeAPI, receives Configuration and CSV file access, Creates the Resty Client
func New(cg config.Config, c CSV) ClientService {
	client := resty.New()
	client.SetHostURL(cg.Client.Host + cg.Client.APIVer)
	return ClientService{client: *client, csv: c}
}

//ImportUser - Applies the bussiness rules to import a new user form a client API into the CSV file
func (c ClientService) ImportUser(id int) (*model.Users, error) {

	resp := request(c, id)
	responseObject := unmarshalResponse(resp)
	err := writeToCSV(responseObject, c)
	if err != nil {
		return nil, errors.New("an error ocurred at saving imported user")
	}
	return &responseObject.Data, nil
}

func writeToCSV(responseObject extUser, c ClientService) error {
	if responseObject.Data.ID != 0 && responseObject.Data.Email != "" {
		err := c.csv.WriteRowData([]string{strconv.Itoa(responseObject.Data.ID), responseObject.Data.Name, responseObject.Data.Email, responseObject.Data.Gender, responseObject.Data.Status})
		if err != nil {
			log.Print(err.Error())
			return err
		}
	}
	return nil
}

func unmarshalResponse(bodyBytes []byte) extUser {

	var responseObject extUser
	err := json.Unmarshal(bodyBytes, &responseObject)
	if err != nil {
		fmt.Print(err.Error())
		return extUser{}
	}
	//fmt.Printf("API Response as struct %+v\n", responseObject)

	return responseObject
}

func request(c ClientService, id int) []byte {
	resp, err := c.client.R().
		SetPathParam("id", strconv.Itoa(id)).
		SetHeader("Accept", "application/json").
		SetHeader("Content-Type", "application/json").
		Get("/users/{id}")

	if err != nil {
		fmt.Print(err.Error())
		return []byte{}
	}

	bodyBytes := resp.Body()

	return bodyBytes
}

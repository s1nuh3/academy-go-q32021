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

//ClientMyAPI - Struc to receibe Client conf
type ClientMyAPI struct {
	client resty.Client
	csv    CSV
}

type CSV interface {
	WriteALLData(records [][]string) error
	WriteRowData(record []string) error
}

type extUser struct {
	Meta interface{} `json:"meta"`
	Data model.Users `json:"data"`
}

//NewClient - Creates the implementation for UseCase ConsumeAPI
func New(cg config.Config, c CSV) ClientMyAPI {
	client := resty.New()
	client.SetHostURL(cg.Client.Host + cg.Client.APIVer)
	return ClientMyAPI{client: *client, csv: c}
}

func (cm ClientMyAPI) ImportUser(id int) (*model.Users, error) {

	resp := request(cm, id)
	responseObject := unmarshalResponse(resp)
	err := writeToCSV(responseObject, cm)
	if err != nil {
		return nil, errors.New("an error ocurred at saving imported user")
	}
	return &responseObject.Data, nil
}

func writeToCSV(responseObject extUser, cm ClientMyAPI) error {
	if responseObject.Data.ID != 0 && responseObject.Data.Email != "" {
		err := cm.csv.WriteRowData([]string{strconv.Itoa(responseObject.Data.ID), responseObject.Data.Name, responseObject.Data.Email, responseObject.Data.Gender, responseObject.Data.Status})
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

func request(cm ClientMyAPI, id int) []byte {
	resp, err := cm.client.R().
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

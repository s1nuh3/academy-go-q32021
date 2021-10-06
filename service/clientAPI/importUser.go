package clientAPI

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/s1nuh3/academy-go-q32021/config"
	"github.com/s1nuh3/academy-go-q32021/model"
)

//ClientMyAPI - Struc to receibe Client conf
type ClientMyAPI struct {
	host string
	csv  CSV
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
	return ClientMyAPI{host: cg.Client.Host + cg.Client.APIVer, csv: c}
}

func (cm ClientMyAPI) ImportUser(id int) (*model.Users, error) {

	resp := request(cm, id)
	responseObject := unmarshalResponse(resp)
	err := WriteToCSV(responseObject, cm)
	if err != nil {
		return nil, errors.New("an error ocurred at saving imported user")
	}
	return &responseObject.Data, nil
}

func WriteToCSV(responseObject extUser, cm ClientMyAPI) error {
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
	cl := &http.Client{}
	//fmt.Println("Host: ", cm.host)
	req, err := http.NewRequest("GET", cm.host+"/users/"+strconv.Itoa(id), nil)
	if err != nil {
		fmt.Print(err.Error())
	}
	//fmt.Printf("Req: %v\n", req)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	resp, err := cl.Do(req)
	if err != nil {
		fmt.Print(err.Error())
		return []byte{}
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Print(err.Error())
		return []byte{}
	}

	return bodyBytes
}

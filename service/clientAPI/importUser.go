package clientAPI

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/s1nuh3/academy-go-q32021/common"
	"github.com/s1nuh3/academy-go-q32021/config"
	"github.com/s1nuh3/academy-go-q32021/model"
)

//ClientMyAPI - Struc to receibe Client conf
type ClientMyAPI struct {
	host string
}

type extUser struct {
	Meta interface{} `json:"meta"`
	Data model.Users `json:"data"`
}

var cfg config.Config

//NewClient - Creates the implementation for UseCase ConsumeAPI
func New(cg config.Config) ClientMyAPI {
	cfg = cg
	return ClientMyAPI{host: cg.Client.Host + cg.Client.APIVer}
}

func (cm ClientMyAPI) ImportUser(id int) (*model.Users, error) {

	resp := request(cm, id)
	//fmt.Printf("API Response Body %+v\n", resp.Body)
	responseObject := unmarshalResponse(resp)

	err := common.WriteRowData(cfg.Csv.Path+cfg.Csv.Name, []string{strconv.Itoa(responseObject.Data.ID), responseObject.Data.Name, responseObject.Data.Email, responseObject.Data.Gender, responseObject.Data.Status})
	if err != nil {
		fmt.Print(err.Error())
		return nil, err
	}

	return &responseObject.Data, nil
}

func unmarshalResponse(bodyBytes []byte) extUser {

	var responseObject extUser
	err := json.Unmarshal(bodyBytes, &responseObject)
	if err != nil {
		fmt.Print(err.Error())
		return extUser{}
	}
	fmt.Printf("API Response as struct %+v\n", responseObject)

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
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Print(err.Error())
		return nil
	}

	return bodyBytes
}

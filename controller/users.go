package controller

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/s1nuh3/academy-go-q32021/model"

	"github.com/gorilla/mux"
)

//UseCaseUser - Interface to be implemented on usecase layer
type UseCaseUser interface {
	GetUser(id int) (*model.Users, error)
	ListUsers() (*[]model.Users, error)
}

//UserHandler - Struc to implement the user handlers
type UserHandler struct {
	ucu UseCaseUser
}

//NewUser - Creates a new instance of handlers for the user paths
func NewUser(ucu UseCaseUser) UserHandler {
	return UserHandler{ucu}

}

// GetUsers - Returns the list of users
func (uh UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {

	u, err := uh.ucu.ListUsers()
	if err != nil {
		returnError(w, r, err, 500)
	}
	w.Header().Set("Content-Type", "application/json")
	if len(*u) != 0 {
		err = json.NewEncoder(w).Encode(u)
		if err != nil {
			returnError(w, r, err, 500)
		}
		return
	} else {
		jso, err := json.Marshal([]int{})
		if err != nil {
			returnError(w, r, err, 500)
		}
		w.Write(jso)
	}
}

// GetUsersbyId - Look up for a user id
func (c UserHandler) GetUsersbyId(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		returnError(w, r, errors.New("ID provided is not valid"), 400)
		return
	}
	u, err := c.ucu.GetUser(id)
	if err != nil {
		returnError(w, r, err, 500)
		return
	}
	if u.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
	} else {
		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(u)
		if err != nil {
			returnError(w, r, err, 500)
		}
	}

}

//ConcurrentRead - Read concurrently the users file
func (c UserHandler) ConcurrentRead(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	filterType := vars["type"]
	items := vars["items"]
	items_per_worker := vars["items_per_worker"]
	//fmt.Println("Type: " + filterType + " Items: " + items + " Items x Worker: " + items_per_worker)

	if filterType == "" {
		returnError(w, r, errors.New("type provided is not valid"), 400)
		return
	}
	i, err := strconv.Atoi(items)
	if err != nil {
		i = 100
	}
	ipw, err := strconv.Atoi(items_per_worker)
	if err != nil {
		ipw = 50
	}

	if ipw > i {
		returnError(w, r, errors.New("items per worker can't be higher than items"), 400)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(filterType + strconv.Itoa(i) + strconv.Itoa(ipw))
	if err != nil {
		returnError(w, r, err, 500)
	}
}

//IndexHandler - Handles the calls to the root path of the server
func (c UserHandler) IndexHandler(w http.ResponseWriter, r *http.Request) {
	err := json.NewEncoder(w).Encode("Welcome, this Go Rest API is to fullfill the Wizeline Academy Go Bootcamp!!")
	if err != nil {
		returnError(w, r, err, 500)
	}
}

func returnError(w http.ResponseWriter, r *http.Request, err error, status int) {
	log.Println(err.Error())
	http.Error(w, err.Error(), status)
}

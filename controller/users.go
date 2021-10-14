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

//UseCaseGoRoutines - Interface to be implemented on usecase layer

//UserHandler - Struc to implement the user handlers
type UserHandler struct {
	ucu UseCaseUser
}

//NewUser - Creates a new instance of handlers for the user paths
func NewUser(ucu UseCaseUser) UserHandler {
	return UserHandler{ucu}
}

// GetUsers - Returns the list of users
func (uh UserHandler) GetUsersHdl(w http.ResponseWriter, r *http.Request) {
	u, err := uh.ucu.ListUsers()
	if err != nil {
		returnError(w, r, err, http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	if len(*u) != 0 {
		err = json.NewEncoder(w).Encode(u)
		if err != nil {
			returnError(w, r, err, http.StatusInternalServerError)
		}
		return
	} else {
		jso, err := json.Marshal([]int{})
		if err != nil {
			returnError(w, r, err, http.StatusInternalServerError)
		}
		w.Write(jso)
	}
}

// GetUsersbyId - Look up for a user id
func (c UserHandler) GetUsersbyIdHdl(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		returnError(w, r, errors.New("ID provided is not valid"), http.StatusBadRequest)
		return
	}
	u, err := c.ucu.GetUser(id)
	if err != nil {
		returnError(w, r, err, http.StatusInternalServerError)
		return
	}
	if u.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
	} else {
		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(u)
		if err != nil {
			returnError(w, r, err, http.StatusInternalServerError)
		}
	}
}

//IndexHandler - Handles the calls to the root path of the server
func (c UserHandler) IndexHdl(w http.ResponseWriter, r *http.Request) {
	err := json.NewEncoder(w).Encode("Welcome, this Go Rest API is to fullfill the Wizeline Academy Go Bootcamp!!")
	if err != nil {
		returnError(w, r, err, http.StatusInternalServerError)
	}
}

func returnError(w http.ResponseWriter, r *http.Request, err error, status int) {
	log.Println(err.Error())
	http.Error(w, err.Error(), status)
}

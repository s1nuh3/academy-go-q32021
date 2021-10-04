package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/s1nuh3/academy-go-q32021/services"

	"github.com/gorilla/mux"
)

// GetUsers - Returns the list of users
func GetUsers(s *services.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(s)
		u, err := s.ListUsers()
		if err != nil {
			returnError(w, r, err, 500)
		}
		w.Header().Set("Content-Type", "application/json")
		if len(*u) != 0 {
			json.NewEncoder(w).Encode(u)
			return
		} else {
			jso, _ := json.Marshal([]int{})
			w.Write(jso)
		}
	}
}

// GetUsersbyId - Look up for a user id
func GetUsersbyId(s *services.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			returnError(w, r, errors.New("ID provided is not valid"), 400)
			return
		}
		u, err := s.GetUser(id)
		if err != nil {
			returnError(w, r, err, 500)
			return
		}
		if u.ID == 0 {
			w.WriteHeader(http.StatusNotFound)
		} else {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(u)
		}

	}
}

func returnError(w http.ResponseWriter, r *http.Request, err error, status int) {
	http.Error(w, err.Error(), status)
}

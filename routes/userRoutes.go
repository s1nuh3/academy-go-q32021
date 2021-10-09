package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

//UserController - Interface to be implemented on controller layer
type UserController interface {
	GetUsers(w http.ResponseWriter, r *http.Request)
	GetUsersbyId(w http.ResponseWriter, r *http.Request)
	IndexHandler(w http.ResponseWriter, r *http.Request)
	ConcurrentRead(w http.ResponseWriter, r *http.Request)
}

//ImportController - Interface to be implemented on controller layer
type ImportController interface {
	ImportHandler(w http.ResponseWriter, r *http.Request)
}

//New - Creates a new instance of routes from controllers handlers
func New(uc UserController, ic ImportController) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", uc.IndexHandler).Methods("GET")
	r.HandleFunc("/users", uc.GetUsers).Methods("GET")
	r.HandleFunc("/users/{id}", uc.GetUsersbyId).Methods("GET")
	r.HandleFunc("/users/import/{id}", ic.ImportHandler).Methods("GET")
	r.HandleFunc("/users/read/", uc.ConcurrentRead).
		Queries("type", "{type:(?:odd|even)}", "items", "{items:[0-9]+}", "items_per_worker", "{items_per_worker:[0-9]+}").Methods("GET")
	r.HandleFunc("/users/read/", uc.ConcurrentRead).
		Queries("type", "{type:(?:odd|even)}", "items", "{items:[0-9]+}").Methods("GET")
	r.HandleFunc("/users/read/", uc.ConcurrentRead).
		Queries("type", "{type:(?:odd|even)}").Methods("GET")
	r.HandleFunc("/users/read/", uc.ConcurrentRead).Methods("GET")
	return r
}

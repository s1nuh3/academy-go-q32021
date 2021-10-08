package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Controller interface {
	GetUsers(w http.ResponseWriter, r *http.Request)
	GetUsersbyId(w http.ResponseWriter, r *http.Request)
	IndexHandler(w http.ResponseWriter, r *http.Request)
	ConcurrentRead(w http.ResponseWriter, r *http.Request)
}

type RouteImportInterface interface {
	ImportUserRte(w http.ResponseWriter, r *http.Request)
}

func New(c Controller, ci RouteImportInterface) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", c.IndexHandler).Methods("GET")
	r.HandleFunc("/users", c.GetUsers).Methods("GET")
	r.HandleFunc("/users/{id}", c.GetUsersbyId).Methods("GET")
	r.HandleFunc("/users/import/{id}", ci.ImportUserRte).Methods("GET")
	r.HandleFunc("/users/read/", c.ConcurrentRead).
		Queries("type", "{type:(?:odd|even)}", "items", "{items:[0-9]+}", "items_per_worker", "{items_per_worker:[0-9]+}")
	return r
}

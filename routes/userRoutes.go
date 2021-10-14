package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

//UserController - Interface to be implemented on controller layer
type UserController interface {
	GetUsersHdl(w http.ResponseWriter, r *http.Request)
	GetUsersbyIdHdl(w http.ResponseWriter, r *http.Request)
	IndexHdl(w http.ResponseWriter, r *http.Request)
}

//ImportController - Interface to be implemented on controller layer
type ImportController interface {
	ImportHdl(w http.ResponseWriter, r *http.Request)
}

//ImportController - Interface to be implemented on controller layer
type GoRoutineContoller interface {
	GoRoutineHdl(w http.ResponseWriter, r *http.Request)
}

//New - Creates a new instance of routes from controllers handlers
func New(uc UserController, ic ImportController, gr GoRoutineContoller) *mux.Router {
	//todo
	r := mux.NewRouter()
	r.HandleFunc("/", uc.IndexHdl).Methods(http.MethodGet)
	r.HandleFunc("/users", uc.GetUsersHdl).Methods(http.MethodGet)
	r.HandleFunc("/users/{id}", uc.GetUsersbyIdHdl).Methods(http.MethodGet)
	r.HandleFunc("/users/import/{id}", ic.ImportHdl).Methods(http.MethodGet)
	r.HandleFunc("/users/read/", gr.GoRoutineHdl).
		Queries("type", "{type:(?:odd|even)}", "items", "{items:[0-9]+}", "items_per_worker", "{items_per_worker:[0-9]+}").Methods(http.MethodGet)
	r.HandleFunc("/users/read/", gr.GoRoutineHdl).
		Queries("type", "{type:(?:odd|even)}", "items", "{items:[0-9]+}").Methods(http.MethodGet)
	r.HandleFunc("/users/read/", gr.GoRoutineHdl).
		Queries("type", "{type:(?:odd|even)}").Methods(http.MethodGet)
	r.HandleFunc("/users/read/", gr.GoRoutineHdl).Methods(http.MethodGet)
	return r
}

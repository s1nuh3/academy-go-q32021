package routes

import "github.com/gorilla/mux"

type Router interface {
	Handlers() *mux.Router
}

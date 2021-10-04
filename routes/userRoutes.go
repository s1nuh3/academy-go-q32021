package routes

import (
	"github.com/s1nuh3/academy-go-q32021/controllers"
	"github.com/s1nuh3/academy-go-q32021/services"

	"github.com/gorilla/mux"
)

type UserRoutes struct {
	us *services.UserService
}

func NewRouter(us *services.UserService) *UserRoutes {
	return &UserRoutes{
		us: us,
	}
}

// Get - Add the handlers to the get methods aviable
func (rs UserRoutes) Handlers() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.IndexHandler()).Methods("GET")
	r.HandleFunc("/users", controllers.GetUsers(rs.us)).Methods("GET")
	r.HandleFunc("/users/{id}", controllers.GetUsersbyId(rs.us)).Methods("GET")
	return r
}

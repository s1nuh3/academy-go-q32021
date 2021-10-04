package app

import (
	"github.com/s1nuh3/academy-go-q32021/repository"
	"github.com/s1nuh3/academy-go-q32021/routes"
	"github.com/s1nuh3/academy-go-q32021/services"

	"github.com/gorilla/mux"
)

// App - This struc is to implement router and other dependencies each time app gets created
type App struct {
	Router *mux.Router
}

// New - Creates a new app that implements routing
func New() *App {
	usersRepo := repository.NewUserCSV("./repository/files/usersdata.csv")
	userUsecase := services.NewService(usersRepo)
	r := routes.NewRouter(userUsecase)
	a := &App{
		Router: r.Handlers(),
	}
	return a
}

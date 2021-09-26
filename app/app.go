package app

import (
	"github.com/s1nuh3/academy-go-q32021/routes"

	"github.com/gorilla/mux"
)

// App - This struc is to implement router and other dependencies each time app gets created
type App struct {
	Router *mux.Router
}

// New - Creates a new app that implements routing
func New() *App {
	A := &App{
		Router: mux.NewRouter(),
	}
	A.initRoutes()
	return A
}

func (a *App) initRoutes() {
	routes.Get(a.Router)
}

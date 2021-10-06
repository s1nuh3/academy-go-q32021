package main

import (
	"log"
	"net/http"

	"github.com/s1nuh3/academy-go-q32021/config"
	"github.com/s1nuh3/academy-go-q32021/controller"
	"github.com/s1nuh3/academy-go-q32021/repository"
	"github.com/s1nuh3/academy-go-q32021/routes"
	"github.com/s1nuh3/academy-go-q32021/service/clientAPI"
	"github.com/s1nuh3/academy-go-q32021/service/user"
	"github.com/s1nuh3/academy-go-q32021/usecase"
)

func main() {

	cfg := config.ReadConfig()

	log.Println("Starting")
	file := repository.New(cfg.Csv.Path + cfg.Csv.Name)
	userService := user.New(file)
	uc := usecase.NewUser(userService)
	c := controller.New(uc)

	client := clientAPI.New(cfg, file)
	iu := usecase.NewConsumeAPI(client, userService)
	ci := controller.NewImportCtrl(iu)

	r := routes.New(c, ci)
	port := cfg.Server.Port
	http.HandleFunc("/", r.ServeHTTP)

	log.Println("App running.. on port ", port)
	log.Fatal(http.ListenAndServe(port, nil))
}

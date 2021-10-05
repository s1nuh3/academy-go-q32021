package main

import (
	"log"
	"net/http"
	"os"

	"github.com/s1nuh3/academy-go-q32021/config"
	"github.com/s1nuh3/academy-go-q32021/controller"
	"github.com/s1nuh3/academy-go-q32021/routes"
	"github.com/s1nuh3/academy-go-q32021/service/clientAPI"
	"github.com/s1nuh3/academy-go-q32021/service/storage"
	"github.com/s1nuh3/academy-go-q32021/usecase"
)

func main() {

	cfg := config.ReadConfig()

	log.Println("Starting")

	repo := storage.New(cfg.Csv.Path + cfg.Csv.Name)
	u := usecase.New(repo)
	c := controller.New(u)

	client := clientAPI.New(cfg)
	iu := usecase.NewConsumeAPI(client)
	ci := controller.NewImportCtrl(iu)

	r := routes.New(c, ci)

	port := cfg.Server.Port
	http.HandleFunc("/", r.ServeHTTP)

	log.Println("App running.. on port ", port)
	err := http.ListenAndServe(port, nil)
	errHandler(err)
}

func errHandler(e error) {
	if e != nil {
		log.Println(e)
		os.Exit(1)
	}
}

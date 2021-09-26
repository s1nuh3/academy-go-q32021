package main

import (
	"log"
	"net/http"
	"os"

	"github.com/s1nuh3/academy-go-q32021/app"
)

func main() {
	app := app.New()
	port := ":8889"
	http.HandleFunc("/", app.Router.ServeHTTP)

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

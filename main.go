package main

import (
	"log"
	"net/http"
	"os"

	"github.com/s1nuh3/academy-go-q32021/config"
	"github.com/s1nuh3/academy-go-q32021/controller"
	"github.com/s1nuh3/academy-go-q32021/repository"
	"github.com/s1nuh3/academy-go-q32021/routes"
	"github.com/s1nuh3/academy-go-q32021/service/clientAPI"
	"github.com/s1nuh3/academy-go-q32021/service/user"
	"github.com/s1nuh3/academy-go-q32021/service/workerpool"
	"github.com/s1nuh3/academy-go-q32021/usecase"
)

func main() {

	cfg := config.ReadConfig()
	cvs := OpenFile(cfg.Csv.Path + cfg.Csv.Name)
	repo := repository.New(cvs)
	userService := user.New(repo)
	userUseCase := usecase.NewUser(userService)
	userHandlers := controller.NewUser(userUseCase)

	client := clientAPI.New(cfg, repo)
	imporUserUseCase := usecase.NewImportUser(client, userService)
	importHandlers := controller.NewImportHandler(imporUserUseCase)

	goRoutineSrv := workerpool.New(cvs)
	goRoutineUseCase := usecase.NewGoRoutine(goRoutineSrv)
	goRoutineHandler := controller.NewGoRoutine(goRoutineUseCase)

	r := routes.New(userHandlers, importHandlers, goRoutineHandler)
	port := cfg.Server.Port
	http.HandleFunc("/", r.ServeHTTP)

	log.Println("App running.. on port ", port)

	log.Fatal(http.ListenAndServe(port, nil))
}

// OpenFile - Reads file from a given path, returns it os.file or error
func OpenFile(filename string) *os.File {
	file, err := os.OpenFile(filename, os.O_RDWR, 0644)
	if err != nil {
		log.Fatal(err.Error())
	}
	//defer file.Close()
	return file
}

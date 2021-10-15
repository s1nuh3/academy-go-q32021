package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

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
	file := openFile(cfg.Csv.Path + cfg.Csv.Name)
	repo := repository.New(file)
	userService := user.New(repo)
	userUseCase := usecase.NewUser(userService)
	userHandler := controller.NewUser(userUseCase)

	clientSrv := clientAPI.New(cfg, repo)
	imporUserUseCase := usecase.NewImportUser(clientSrv, userService)
	importHandler := controller.NewImportHandler(imporUserUseCase)

	goRoutineSrv := workerpool.New(file)
	goRoutineUseCase := usecase.NewGoRoutine(goRoutineSrv)
	goRoutineHandler := controller.NewGoRoutine(goRoutineUseCase)

	r := routes.New(userHandler, importHandler, goRoutineHandler)
	port := cfg.Server.Port

	srv := &http.Server{
		Addr:    port,
		Handler: r,
	}

	log.Println("HTTP server listening on port", port)

	idleConnsClosed := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt, syscall.SIGINT)
		<-sigint
		log.Println("HTTP server Shutdown...")
		if err := srv.Shutdown(context.Background()); err != nil {
			log.Printf(" HTTP server Shutdown: %v", err)
		}
		close(idleConnsClosed)
	}()

	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		log.Printf("Error HTTP server ListenAndServe: %v", err)
	}
	<-idleConnsClosed
}

// openFile - Reads file from a given path, returns it os.file or error
func openFile(filename string) *os.File {
	file, err := os.OpenFile(filename, os.O_RDWR, 0644)
	if err != nil {
		log.Fatal(err.Error())
	}
	return file
}

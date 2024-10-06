package main

import (
	"go-todo/internal/handler"
	"go-todo/internal/repository"
	"go-todo/internal/service"
	"go-todo/pkg/server"
	"log"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(server.Server)
	if err := srv.Run("8000", handlers.InitRouter()); err != nil {
		log.Fatalf("failed to run server: %s", err.Error())
	}
}

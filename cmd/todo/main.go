package main

import (
	"go-todo/pkg/server"
	"log"

	routerHandler "go-todo/internal/handler"
)

func main() {
	handler := new(routerHandler.Handler)
	srv := new(server.Server)

	if err := srv.Run("8000", handler.InitRouter()); err != nil {
		log.Fatalf("failed to run server: %s", err.Error())
	}
}

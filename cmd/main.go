package main

import (
	"github.com/XXena/todo-app"
	"github.com/XXena/todo-app/package/handler"
	"github.com/XXena/todo-app/package/repository"
	"github.com/XXena/todo-app/package/service"
	"log"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error ocurred while running http server: %s", err.Error())
	}
}

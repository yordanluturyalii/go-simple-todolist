package main

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/yordanluturyalii/go-simple-todolist/app"
	"github.com/yordanluturyalii/go-simple-todolist/config"
	"github.com/yordanluturyalii/go-simple-todolist/controller"
	"github.com/yordanluturyalii/go-simple-todolist/database"
	"github.com/yordanluturyalii/go-simple-todolist/repository"
	"github.com/yordanluturyalii/go-simple-todolist/service"
	"net/http"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	cnf := config.NewConfig()
	db := database.NewDatabase(cnf)
	validate := validator.New()

	todolistRepository := repository.NewTodolistRepository()
	todolistService := service.NewTodolistService(todolistRepository, validate, db)
	todolistController := controller.NewTodolistController(todolistService)

	router := app.NewRouter(todolistController)

	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

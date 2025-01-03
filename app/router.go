package app

import (
	"github.com/julienschmidt/httprouter"
	"github.com/yordanluturyalii/go-simple-todolist/controller"
)

func NewRouter(todolistController controller.TodolistController) *httprouter.Router {
	router := httprouter.New()

	router.GET("/api", todolistController.FindAll)
	router.GET("/api/:id", todolistController.FindById)

	return router
}

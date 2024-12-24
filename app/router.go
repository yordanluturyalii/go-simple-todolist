package app

import (
	"github.com/julienschmidt/httprouter"
	"github.com/yordanluturyalii/go-simple-todolist/controller"
)

func NewRouter(todolistController controller.TodolistController) *httprouter.Router {
	router := httprouter.New()

	router.GET("/", todolistController.FindAll)

	return router
}

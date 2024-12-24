package controller

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"github.com/yordanluturyalii/go-simple-todolist/model"
	"github.com/yordanluturyalii/go-simple-todolist/service"
	"net/http"
)

type TodolistControllerImpl struct {
	TodolistService service.TodolistService
}

func NewTodolistController(todolistService service.TodolistService) *TodolistControllerImpl {
	return &TodolistControllerImpl{
		TodolistService: todolistService,
	}
}

func (t TodolistControllerImpl) Create(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	//TODO implement me
	panic("implement me")
}

func (t TodolistControllerImpl) Update(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	//TODO implement me
	panic("implement me")
}

func (t TodolistControllerImpl) Delete(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	//TODO implement me
	panic("implement me")
}

func (t TodolistControllerImpl) FindById(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	//TODO implement me
	panic("implement me")
}

func (t TodolistControllerImpl) FindAll(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	todolists := t.TodolistService.FindAll(r.Context())
	response := model.GlobalResponse{
		Message: "Success Get Data",
		Data:    todolists,
		Errors:  nil,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

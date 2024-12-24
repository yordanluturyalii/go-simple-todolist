package controller

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"github.com/yordanluturyalii/go-simple-todolist/helper"
	"github.com/yordanluturyalii/go-simple-todolist/model"
	"github.com/yordanluturyalii/go-simple-todolist/service"
	"net/http"
	"strconv"
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
	var param string = params.ByName("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		helper.PanicIfNil(err)
	}

	var response model.GlobalResponse
	var errResponse model.ErrorResponse

	todolist, err := t.TodolistService.FindById(r.Context(), int64(id))
	if err != nil {
		errResponse = model.ErrorResponse{
			Message: err.Error(),
			Error:   err,
		}

		response = model.GlobalResponse{
			Message: "Failed To Get Data",
			Data:    nil,
			Errors:  errResponse,
		}
	} else {
		response = model.GlobalResponse{
			Message: "Success Get Data By Id",
			Data:    todolist,
			Errors:  nil,
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
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

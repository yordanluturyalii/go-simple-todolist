package service

import (
	"context"
	"database/sql"
	"github.com/go-playground/validator/v10"
	"github.com/yordanluturyalii/go-simple-todolist/helper"
	"github.com/yordanluturyalii/go-simple-todolist/model"
	"github.com/yordanluturyalii/go-simple-todolist/repository"
)

type TodolistServiceImpl struct {
	repository.TodolistRepository
	*validator.Validate
	*sql.DB
}

func NewTodolistService(todolistRepository repository.TodolistRepository, validate *validator.Validate, tx *sql.DB) *TodolistServiceImpl {
	return &TodolistServiceImpl{todolistRepository, validate, tx}
}

func (service *TodolistServiceImpl) Create(ctx context.Context, request model.TodolistRequest) (*model.TodolistResponse, error) {
	err := service.Validate.Struct(request)
	if err != nil {
		return nil, err
	}

	tx, err := service.DB.Begin()
	if err != nil {
		return nil, err
	}
	defer helper.CommitOrRollback(tx)

	todolist := model.Todolist{
		Task: request.Task,
	}

	todolist = service.Save(ctx, service.DB, todolist)

	response := model.TodolistResponse{
		Id:   todolist.Id,
		Task: todolist.Task,
	}

	return &response, nil
}

func (service *TodolistServiceImpl) Update(ctx context.Context, request model.TodolistRequest) (*model.TodolistResponse, error) {
	err := service.Validate.Struct(request)
	if err != nil {
		return nil, err
	}

	tx, err := service.DB.Begin()
	if err != nil {
		return nil, err
	}
	defer helper.CommitOrRollback(tx)

	todolist := model.Todolist{
		Id: request.Id,
	}

	todolist = service.TodolistRepository.Update(ctx, service.DB, todolist)

	response := model.TodolistResponse{
		Id:   todolist.Id,
		Task: todolist.Task,
	}

	return &response, nil
}

func (service *TodolistServiceImpl) Delete(ctx context.Context, id int64) {

	todolist := model.Todolist{Id: id}

	service.TodolistRepository.Delete(ctx, service.DB, todolist)
}

func (service *TodolistServiceImpl) FindById(ctx context.Context, id int64) (*model.TodolistResponse, error) {
	todolist := model.Todolist{Id: id}
	findById, err := service.TodolistRepository.FindById(ctx, service.DB, todolist)
	if err != nil {
		return nil, err
	}

	response := &model.TodolistResponse{
		Id:   findById.Id,
		Task: findById.Task,
	}

	return response, nil
}

func (service *TodolistServiceImpl) FindAll(ctx context.Context) []model.TodolistResponse {
	todolist := model.Todolist{}
	repositories := service.TodolistRepository.FindAll(ctx, service.DB, todolist)

	var todolists []model.TodolistResponse
	for _, repository := range repositories {
		todolists = append(todolists, model.TodolistResponse{Id: repository.Id, Task: repository.Task})
	}

	return todolists
}

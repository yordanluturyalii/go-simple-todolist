package service

import (
	"context"
	"github.com/yordanluturyalii/go-simple-todolist/model"
)

type TodolistService interface {
	Create(ctx context.Context, request model.TodolistRequest) (*model.TodolistResponse, error)
	Update(ctx context.Context, request model.TodolistRequest) (*model.TodolistResponse, error)
	Delete(ctx context.Context, id int64)
	FindById(ctx context.Context, id int64) (*model.TodolistResponse, error)
	FindAll(ctx context.Context) []model.TodolistResponse
}

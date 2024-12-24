package repository

import (
	"context"
	"database/sql"
	"github.com/yordanluturyalii/go-simple-todolist/model"
)

type TodolistRepository interface {
	Save(ctx context.Context, tx *sql.DB, todolist model.Todolist) model.Todolist
	Update(ctx context.Context, tx *sql.DB, todolist model.Todolist) model.Todolist
	Delete(ctx context.Context, tx *sql.DB, todolist model.Todolist)
	FindById(ctx context.Context, tx *sql.DB, todolist model.Todolist) (model.Todolist, error)
	FindAll(ctx context.Context, tx *sql.DB, todolist model.Todolist) []model.Todolist
}

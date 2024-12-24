package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/yordanluturyalii/go-simple-todolist/helper"
	"github.com/yordanluturyalii/go-simple-todolist/model"
)

type TodolistRepositoryImpl struct {
}

func NewTodolistRepository() *TodolistRepositoryImpl {
	return &TodolistRepositoryImpl{}
}

func (t TodolistRepositoryImpl) Save(ctx context.Context, tx *sql.DB, todolist model.Todolist) model.Todolist {
	script := "INSERT INTO todolists (task) values (?)"
	result, err := tx.ExecContext(ctx, script, todolist.Task)
	helper.PanicIfNil(err)

	id, err := result.LastInsertId()
	helper.PanicIfNil(err)

	todolist.Id = id
	return todolist
}

func (t TodolistRepositoryImpl) Update(ctx context.Context, tx *sql.DB, todolist model.Todolist) model.Todolist {
	script := "UPDATE todolists SET task = ? WHERE id = ?"
	_, err := tx.ExecContext(ctx, script, todolist.Task, todolist.Id)
	helper.PanicIfNil(err)

	return todolist
}

func (t TodolistRepositoryImpl) Delete(ctx context.Context, tx *sql.DB, todolist model.Todolist) {
	script := "DELETE FROM todolists WHERE id = ?"
	_, err := tx.ExecContext(ctx, script, todolist.Id)
	helper.PanicIfNil(err)
}

func (t TodolistRepositoryImpl) FindById(ctx context.Context, tx *sql.DB, todolist model.Todolist) (model.Todolist, error) {
	query := "SELECT id, task FROM todolists WHERE id = ?"
	row, err := tx.QueryContext(ctx, query, todolist.Id)
	helper.PanicIfNil(err)
	defer row.Close()

	if row.Next() {
		err := row.Scan(&todolist.Id, &todolist.Task)
		helper.PanicIfNil(err)
		return todolist, nil
	} else {
		return todolist, errors.New("todolist not found")
	}
}

func (t TodolistRepositoryImpl) FindAll(ctx context.Context, tx *sql.DB, todolist model.Todolist) []model.Todolist {
	query := "SELECT id, task FROM todolists"
	rows, err := tx.QueryContext(ctx, query)
	helper.PanicIfNil(err)
	defer rows.Close()

	var todolists []model.Todolist
	for rows.Next() {
		err := rows.Scan(&todolist.Id, &todolist.Task)
		helper.PanicIfNil(err)
		todolists = append(todolists, todolist)
	}

	return todolists
}

package test

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/yordanluturyalii/go-simple-todolist/database"
	"github.com/yordanluturyalii/go-simple-todolist/helper"
	"github.com/yordanluturyalii/go-simple-todolist/model"
	"github.com/yordanluturyalii/go-simple-todolist/repository"
	"testing"
)

func TestSaveSuccess(t *testing.T) {
	conn := database.NewDatabase()
	defer conn.Close()

	repo := repository.NewTodolistRepository()
	ctx := context.Background()

	data := model.Todolist{
		Task: "Walawe",
	}
	result := repo.Save(ctx, conn, data)

	assert.Equal(t, "Walawe", result.Task)
}

func TestUpdateSuccess(t *testing.T) {
	conn := database.NewDatabase()
	defer conn.Close()

	repo := repository.NewTodolistRepository()
	ctx := context.Background()

	data := model.Todolist{
		Id:   1,
		Task: "Walawe",
	}
	result := repo.Update(ctx, conn, data)

	assert.Equal(t, "Walawe", result.Task)
}

func TestDeleteSuccess(t *testing.T) {
	conn := database.NewDatabase()
	defer conn.Close()

	repo := repository.NewTodolistRepository()
	ctx := context.Background()

	data := model.Todolist{
		Id:   1,
		Task: "Walawe",
	}
	repo.Delete(ctx, conn, data)
}

func TestFindByIdSuccess(t *testing.T) {
	conn := database.NewDatabase()
	defer conn.Close()
	defer func() {
		if err := recover(); err != nil {
			assert.Equal(t, "todolist not found", err)
		}
	}()

	repo := repository.NewTodolistRepository()
	ctx := context.Background()

	data := model.Todolist{
		Id: 3,
	}
	result, err := repo.FindById(ctx, conn, data)
	helper.PanicIfNil(err)

	assert.Equal(t, "Walawe", result.Task)
}

func TestFindAll(t *testing.T) {
	conn := database.NewDatabase()
	defer conn.Close()

	repo := repository.NewTodolistRepository()
	ctx := context.Background()

	data := model.Todolist{}
	results := repo.FindAll(ctx, conn, data)

	assert.NotNil(t, results)
}

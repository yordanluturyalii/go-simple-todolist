package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/yordanluturyalii/go-simple-todolist/config"
	"github.com/yordanluturyalii/go-simple-todolist/helper"
	"time"
)

func NewDatabase(cnf *config.Config) *sql.DB {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	dbName := cnf.Env.GetString("DB_NAME")
	dbUser := cnf.Env.GetString("DB_USER")
	dbPass := cnf.Env.GetString("DB_PASS")
	dbHost := cnf.Env.GetString("DB_HOST")
	dbPort := cnf.Env.GetString("DB_PORT")

	dsn := dbUser + ":" + dbPass + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName
	db, err := sql.Open("mysql", dsn)
	helper.PanicIfNil(err)

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(time.Hour)

	return db
}

package database

import (
	"database/sql"
	"fmt"
	"os"
)

var (
	name     = os.Getenv("DATABASE_USERNAME")
	password = os.Getenv("DATABASE_PASSWORD")
	database = os.Getenv("DATABASE_NAME")
)

type SqlHandler struct {
	Conn *sql.DB
}

func NewSqlHandler() *SqlHandler {
	DbConnection, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(db:3306)/%s", name, password, database))
	if err != nil {
		panic(err.Error())
	}
	sqlHandler := new(SqlHandler)
	sqlHandler.Conn = DbConnection
	return sqlHandler
}

package persistence

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var (
	db *sql.DB
)

func connect() *sql.DB {
	t, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME")))
	if err != nil {
		panic(err)
	}
	return t
}

func Instance() *sql.DB {
	if db == nil {
		db = connect()
	}
	return db
}

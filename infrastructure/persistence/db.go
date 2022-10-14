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

func init() {
	var err error
	db, err = sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME")))
	if err != nil {
		panic(err)
	}
}

func Instance() *sql.DB {
	return db
}

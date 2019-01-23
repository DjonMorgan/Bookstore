package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func Init() {
	connect()
}

func connect() {
	DBMS := "mysql"
	USER := "root"
	PASS := "3228"
	PROTOCOL := "tcp(0.0.0.0:3306)"
	DBNAME := "book"
	PARAMS := "?parseTime=true"

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + PARAMS

	var err error

	db, err = sql.Open(DBMS, CONNECT)

	if err != nil {
		panic(err.Error())
	}
}

func Manager() *sql.DB {
	return db
}

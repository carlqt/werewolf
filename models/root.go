package models

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

const (
	dbName   = "werewolf"
	host     = "localhost"
	user     = "postgres"
	password = "postgres"
	dbPort   = 5433
)

func InitDB() *sqlx.DB {
	var err error

	dbinfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host,
		dbPort,
		user,
		password,
		dbName,
	)

	db, err = sqlx.Open("postgres", dbinfo)
	if err != nil {
		panic(err)
	}

	return db
}

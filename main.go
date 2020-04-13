package main

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	db_name  = "werewolf"
	host     = "localhost"
	user     = "postgres"
	password = ""
	port     = 5432
)

func main() {

	dbinfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host,
		port,
		user,
		password,
		db_name,
	)

	db, err := sqlx.Open("postgres", dbinfo)
	defer db.Close()

	if err != nil {
		panic(err)
	}

	err = db.Ping()

	if err != nil {
		panic(err)
	}

	fmt.Println("Postgres connected")
}

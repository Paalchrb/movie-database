package config

import (
	"database/sql"
	"fmt"

	//driver for posgres package
	_ "github.com/lib/pq"
)

//DB initializes connection to postgres database
var DB *sql.DB

func init() {
	var err error
	DB, err = sql.Open("postgres", "postgres://pal-christianby:password@localhost/movieapi?sslmode=disable")
	if err != nil {
		panic(err)
	}

	if err = DB.Ping(); err != nil {
		panic(err)
	}
	fmt.Println("You connected to your database.")
}

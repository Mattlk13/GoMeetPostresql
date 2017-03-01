package main

import (
	/* import standard sql package to  connect to db */
	"database/sql"
	"fmt"
	/* use lib/pq as a postgres driver */
	_ "github.com/lib/pq"
)

/* database connection vars */
const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "password"
	dbname   = "postgres"
)

func main() {
	/* connection string */
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	/* validate connection to pg */
	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		panic(err)
	}
	defer db.Close()

	/* open a connection to the database */
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")

}

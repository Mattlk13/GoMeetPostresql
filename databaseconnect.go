package main

import (
	/* import standard sql package to  connect to db */
	"database/sql"
	"fmt"
	/* use reflect to find type of vars */
	"reflect"
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

func query_database(query_string string) *sql.Rows {
	/* connection string */
	connection_string := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	/* validate connection to pg */
	db, err := sql.Open("postgres", connection_string)

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

	returned_output, err := db.Query(query_string)
	if err != nil {
		panic(err)
	}

	fmt.Println(reflect.TypeOf(returned_output))

	return returned_output
}

func main() {
	query := "SELECT * FROM timecards"
	/* connect to database and load into var */
	list_of_timecards := query_database(query)

	fmt.Println(list_of_timecards)

}

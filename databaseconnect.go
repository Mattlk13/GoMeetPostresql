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
		fmt.Println("Error: ", err)
	}
	defer db.Close()

	/* open a connection to the database */
	err = db.Ping()
	if err != nil {
		fmt.Println("Error: ", err)
	}

	fmt.Println("Successfully connected!")

	/* Query Database */
	returned_output, err := db.Query(query_string)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	defer returned_output.Close()

	/* Output result */

	fmt.Println(reflect.TypeOf(returned_output))
	return returned_output
}

func insert_into_database(insert_string string) {

	/* connection string */
	connection_string := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	/* validate connection to pg */
	db, err := sql.Open("postgres", connection_string)

	if err != nil {
		fmt.Println("Error: ", err)
	}
	defer db.Close()

	/* open a connection to the database */
	err = db.Ping()
	if err != nil {
		fmt.Println("Error: ", err)
	}

	fmt.Println("Successfully connected!")

	/* Insert into Database */
	insert_statement, err := db.Prepare(insert_string)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	_, err = insert_statement.Exec()
	if err != nil {
		fmt.Println("Error: ", err)
	}else{
		fmt.Println("Successfully Inserted")
	}
}

func main() {
	query_all := "SELECT * FROM timecards"
	query_one := "SELECT * FROM timecards WHERE(Id = 1)"
	insert_one := "INSERT INTO timecards (ID, USERNAME, OCCURRENCE) VALUES(5, 'RICKY', current_timestamp)"

	/* connect to database and load into var */
	list_of_timecards := query_database(query_all)
	first_timecard := query_database(query_one)

	fmt.Println(list_of_timecards)
	fmt.Println(first_timecard)

	insert_into_database(insert_one)

}

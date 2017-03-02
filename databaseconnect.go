package main

import (
	/* import standard sql package to  connect to db */
	"database/sql"
	"fmt"
	/* use reflect to find type of vars */
	/*"reflect"*/
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

func query_database(db *sql.DB, query_string string) *sql.Rows {
	/* open a connection to the database */
	err := db.Ping()
	if err != nil {
		fmt.Println("Error!")
		panic(err)
	}

	/* Query Database */
	returned_rows, err := db.Query(query_string)
	if err != nil {
		fmt.Println("Error!")
		panic(err)
	}
	defer returned_rows.Close()

	column_names, err := returned_rows.Columns()
	if err != nil {
		fmt.Println("Error!")
		panic(err)
	}

	column_length := len(column_names)

	/* Set up slices for scanning rows*/
	values := make([]interface{}, column_length)
	value_pointers := make([]interface{}, column_length)

	for returned_rows.Next() {
		for i, _ := range column_names {
			value_pointers[i] = &values[i]
		}

		returned_rows.Scan(value_pointers...)
		for i, column := range column_names {
			var v interface{}
			value := values[i]
			byte_rep, ok := value.([]byte)

			if ok {
				v = string(byte_rep)
			} else {
				v = value
			}

			fmt.Println(column, v)
		}
	}

	err = returned_rows.Err()

	if err != nil {
		fmt.Println("Error!")
		panic(err)
	}

	fmt.Println("Successfully retrieved data!")

	/* Output result */
	return returned_rows
}

func insert_to_database(db *sql.DB, insert_statement string) {
	/* open a connection to the database */
	err := db.Ping()
	if err != nil {
		fmt.Println("Error!")
		panic(err)
	}

	/* Prepare Insertion Database */
	insert_command, err := db.Prepare(insert_statement)
	if err != nil {
		fmt.Println("Error!")
		panic(err)
	}

	/* Execute Insert Statement */
	_, err = insert_command.Exec()
	if err != nil {
		fmt.Println("Error!")
		panic(err)
	}

	fmt.Println("Successfully Inserted")

	defer insert_command.Close()
}

func main() {
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

	fmt.Println("Successfully connected to database!")

	/* Query Strings */
	query_one := "SELECT * FROM timecards WHERE(Id = 1)"

	/* Insert Strings */
	insert_statement := "INSERT INTO timecards (ID, USERNAME, OCCURRENCE) " +
		"VALUES(15, 'ROHAN', current_timestamp)"

	/* Database Operations */
	first_timecard := query_database(db, query_one)

	fmt.Println(first_timecard)

	insert_to_database(db, insert_statement)

}

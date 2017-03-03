package main

import (
	/* import buffered i/o for interface */
	"bufio"
	/* import standard sql package to  connect to db */
	"database/sql"
	/* import fmt and reflect for debugging */
	"fmt"
	//"reflect"
	/* import os for user input */
	"os"
	/* import strings for string manipulation */
	"strings"
	/* use lib/pq as a postgres driver */
	_ "github.com/lib/pq"
)

/* EDIT THESE FOR YOUR OWN DATABASE -- Database connection vars */
const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "password"
	dbname   = "postgres"
)

func query_database(db *sql.DB, query_string string) []interface{} {
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

	return output_rows(returned_rows)
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
	new_rows, err := insert_command.Exec()
	if err != nil {
		fmt.Println("Error!")
		panic(err)
	}
	defer insert_command.Close()

	/* Output success message of rows changed */
	number_of_rows_changed, _ := new_rows.RowsAffected()
	db_insert_success_message := fmt.Sprintf("Successfully inserted %d row", number_of_rows_changed)
	fmt.Println(db_insert_success_message)

}

func output_rows(returned_rows *sql.Rows) []interface{} {
	/* Get column names and length */
	column_names, err := returned_rows.Columns()
	if err != nil {
		fmt.Println("Error!")
		panic(err)
	}
	column_length := len(column_names)

	/* output array */
	var query_output []interface{}

	/* Iterate over every row returned from query */
	for returned_rows.Next() {
		/* Set up empty slices to hold values and pointers to values */
		values := make([]interface{}, column_length)
		value_pointers := make([]interface{}, column_length)

		/* Row assembly */
		row_map := make(map[string]interface{})

		/* Iterate over every column name and set pointers to empty values */
		for i, _ := range column_names {
			value_pointers[i] = &values[i]
		}

		/* Use scan to convert returned_rows elements to pointed types, and store */
		returned_rows.Scan(value_pointers...)

		/* Assign value to key in row map */
		for i, column := range column_names {
			row_map[column] = values[i]
		}

		/* Add formatted row_map to output array */
		query_output = append(query_output, row_map)
	}

	/* Output result */
	db_select_success_message := fmt.Sprintf("\nSuccessfully returned %d rows\n", len(query_output))
	fmt.Println(db_select_success_message)

	return query_output
}

func display_table(db *sql.DB, table_name string){
	select_all_query := fmt.Sprintf("SELECT * FROM %s", table_name)
	raw_rows := query_database(db, select_all_query)
	for _, row := range raw_rows {
		fmt.Println(row)
	}
}

func show_help() {
	fmt.Println("\n	Thanks for checking out this little toy. Here's what the wrapper supports:\n")
	fmt.Println("	- SELECT	Select rows from tables 'SELECT * FROM users WHERE(id = 1)'")
	fmt.Println("	- INSERT	Insert row into table 'INSERT INTO users (ID, NAME) VALUES(1, 'Rohan')'")
	fmt.Println("	- SHOW		Display an entire table 'SHOW users'")
	fmt.Println("	- QUIT		Type -q to quit\n")
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
		panic(err)
	}
	defer db.Close()

	db_connect_success_message := fmt.Sprintf("Successfully connected to %s database!", dbname)
	fmt.Println(db_connect_success_message)

	/* Database Operations
	all_timecard := query_database(db, "SELECT * FROM timecards") */

	/* Get one row from database return
	entry := all_timecard[0].(map[string]interface{})
	fmt.Println(entry) */

	/* Insert Strings
	insert_statement := "INSERT INTO timecards (ID, USERNAME, OCCURRENCE) " +
		"VALUES(15, 'ROHAN', current_timestamp)" */

	/* Prompt user to query database */
	input_reader := bufio.NewScanner(os.Stdin)
	var input_string string
	prompt := fmt.Sprintf("%s:", dbname)

	fmt.Println("Enter raw SQL to query database, or '-h' for help")
	for input_string != "-q" {
		fmt.Print(prompt)
		input_reader.Scan()
		input_string = input_reader.Text()
		input_string_split := strings.Split(input_string, " ")
		sql_operation := input_string_split[0]
		sql_operation = strings.ToLower(sql_operation)
		if input_string == "-h" {
			show_help()
		} else if input_string != "-q" {
			switch sql_operation {
			case "select":
				select_output := query_database(db, input_string)
				fmt.Println(select_output)
			case "insert":
				insert_to_database(db, input_string)
			case "show":
				if (len(input_string_split) == 2){
					table_name := input_string_split[1]
					display_table(db, table_name)
				}else {
					fmt.Println("Invalid number of arguments. Try 'Show table_name'")
				}
			default:
				fmt.Println("Operation not supported. Type -h for help")
			}

		} else {
			db.Close()
			fmt.Println("Closed database connection, goodbye :)")
		}
	}
}

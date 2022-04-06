package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3" // Import go-sqlite3 library
	"go-sqlite/database"
	table "go-sqlite/database/tables"
)

func main() {

	fmt.Println("Starting go-sqlite...")

	createSchema()

	populateSchema()

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {

		}
	}(database.GetDb())
}

func createSchema() {
	table.Student{}.Create()
	table.Department{}.Create()
}

func populateSchema() {
	for i := 0; i < 1; i++ {
		table.Department{Department: "IT"}.Insert()
	}
	for i := 0; i < 3; i++ {
		table.Student{Code: i, FirstName: "Justin", LastName: "Mifkovich", Department: i + 1}.Insert()
	}

	table.DisplayStudents()
}

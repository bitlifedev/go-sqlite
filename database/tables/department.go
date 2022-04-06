package table

import (
	"go-sqlite/database"
	"log"
)

type Department struct {
	id         int
	Department string
}

func (d Department) Create() {
	createTableSQL := `CREATE TABLE department (
		"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,		
		"department" Text 	
	  );` // SQL Statement for Create Table

	log.Println("Create department table...")

	statement, err := database.GetDb().Prepare(createTableSQL) // Prepare SQL Statement
	if err != nil {
		log.Println(err.Error())
	}
	statement.Exec() // Execute SQL Statements
	log.Println("department table created")

}

func (d Department) Insert() {

	//log.Println("Inserting student record ...")
	insertSQL := `INSERT INTO department (department) VALUES (?)`
	statement, err := database.GetDb().Prepare(insertSQL) // Prepare statement.
	// This is good to avoid SQL injections
	if err != nil {
		log.Fatalln(err.Error())
	}
	_, err = statement.Exec(d.Department)
	if err != nil {
		log.Println(insertSQL + " " + err.Error())
	} else {
		log.Println("Inserted department")
	}

}

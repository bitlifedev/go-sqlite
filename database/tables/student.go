package table

import (
	"go-sqlite/database"
	"log"
)

type Student struct {
	Code       int
	FirstName  string
	LastName   string
	Department int
}

func (s Student) Create() {
	createTableSQL := `CREATE TABLE student (
		"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,		
		"code" INTEGER ,
		"firstname" TEXT,
		"lastname" TEXT,
		"departmentId" INTEGER,
		FOREIGN KEY("departmentId") REFERENCES department(id) 	
	  );` // SQL Statement for Create Table

	log.Println("Create student table...")

	statement, err := database.GetDb().Prepare(createTableSQL) // Prepare SQL Statement
	if err != nil {
		log.Println(err.Error())
	}
	statement.Exec() // Execute SQL Statements
	log.Println("student table created")

}

func (s Student) Insert() {

	//log.Println("Inserting student record ...")
	insertSQL := `INSERT INTO student(code, firstname, lastname, departmentId) VALUES (?, ?, ?, ?)`
	statement, err := database.GetDb().Prepare(insertSQL) // Prepare statement.
	// This is good to avoid SQL injections
	if err != nil {
		log.Fatalln(err.Error())
	}
	_, err = statement.Exec(s.Code, s.FirstName, s.LastName, s.Department)
	if err != nil {
		log.Println(insertSQL + " " + err.Error())
	} else {
		log.Println("Inserted student")
	}
}

func DisplayStudents() {
	row, err := database.GetDb().Query("SELECT * FROM student ORDER BY id")
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()
	for row.Next() { // Iterate and fetch the records from result cursor
		var id int
		var code int
		var firstname string
		var lastname string
		var department int
		row.Scan(&id, &code, &firstname, &lastname, &department)
		log.Println(code, " ", firstname, " ", lastname, " ", department)
	}
}

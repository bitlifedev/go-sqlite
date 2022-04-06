package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

var (
	db *sql.DB
)

func init() {
	fmt.Println("Starting go-sqlite...")
	os.Remove("sqlite-database.db") // I delete the file to avoid duplicated records.
	// SQLite is a file based database.
	log.Println("Creating sqlite-database.db...")
	file, err := os.Create("sqlite-database.db") // Create SQLite file
	if err != nil {
		log.Fatal(err.Error())
	}
	file.Close()
	log.Println("sqlite-database.db created")

}
func GetDb() *sql.DB {
	db, _ := sql.Open("sqlite3", "./sqlite-database.db?_foreign_keys=on") // Open the created SQLite File
	return db
}

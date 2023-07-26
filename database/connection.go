package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

// Koneksi ke database MySQL
func ConnectDB() {
	var err error
	db, err = sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/esports_management")
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	// Set the maximum number of open connections to the database
	db.SetMaxOpenConns(10)

	// Check the connection to the database
	if err = db.Ping(); err != nil {
		log.Fatal("Failed to ping the database:", err)
	}

	fmt.Println("Connected to the database!")
}

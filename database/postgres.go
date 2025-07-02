package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var Conn *sql.DB

// InitDB initializes the database connection using environment variables.
// It reads the database connection parameters from environment variables and establishes a connection.
func InitDB() {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname,
	)

	var err error
	Conn, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Failed to open DB connection: ", err)
	}

	if err = Conn.Ping(); err != nil {
		log.Fatal("Failed to ping DB: ", err)
	}

	log.Println("Connected to database")
}

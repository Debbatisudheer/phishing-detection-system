package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectDatabase() {

	var connStr string

	databaseURL := os.Getenv("DATABASE_URL")

	if databaseURL != "" {

		connStr = databaseURL
		fmt.Println("Using DATABASE_URL")

	} else {

		host := os.Getenv("DB_HOST")
		if host == "" {
			host = "localhost"
		}

		port := os.Getenv("DB_PORT")
		if port == "" {
			port = "5432"
		}

		user := os.Getenv("DB_USER")
		if user == "" {
			user = "postgres"
		}

		password := os.Getenv("DB_PASSWORD")
		if password == "" {
			password = "sudheer"
		}

		dbname := os.Getenv("DB_NAME")
		if dbname == "" {
			dbname = "phishing_platform"
		}

		connStr = fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			host,
			port,
			user,
			password,
			dbname,
		)

		fmt.Println("Using Local PostgreSQL")
	}

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	if err := db.Ping(); err != nil {
		panic(err)
	}

	fmt.Println("Database Connected Successfully")
	DB = db
}

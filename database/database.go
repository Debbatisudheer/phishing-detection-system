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

	// Railway
	databaseURL := os.Getenv("DATABASE_URL")

	if databaseURL != "" {

		connStr = databaseURL

		fmt.Println(
			"Using Railway DATABASE_URL",
		)

	} else {

		// Local Development

		connStr =
			"user=postgres " +
				"password=sudheer " +
				"dbname=phishing_platform " +
				"sslmode=disable"

		fmt.Println(
			"Using Local PostgreSQL",
		)
	}

	db, err := sql.Open(
		"postgres",
		connStr,
	)

	if err != nil {

		panic(err)

	}

	err = db.Ping()

	if err != nil {

		panic(err)

	}

	fmt.Println(
		"Database Connected Successfully",
	)

	DB = db
}
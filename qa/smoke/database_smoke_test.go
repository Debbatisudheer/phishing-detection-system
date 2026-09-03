package smoke

import (
	"testing"

	"phishing-platform/database"
)

func TestDatabaseSmoke(t *testing.T) {

	setupDatabase()

	if database.DB == nil {

		t.Fatal("Database connection is nil")

	}

	err := database.DB.Ping()

	if err != nil {

		t.Fatalf(
			"Database ping failed: %v",
			err,
		)

	}

	t.Log("==============================")
	t.Log("DATABASE SMOKE TEST PASSED")
	t.Log("Database Connected")
	t.Log("Database Ping Successful")
	t.Log("==============================")
}

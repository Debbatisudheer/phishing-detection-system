package api

import (
	"sync"

	"phishing-platform/database"

	"github.com/joho/godotenv"
)

var once sync.Once

func setupDatabase() {

	once.Do(func() {

		_ = godotenv.Load()

		database.ConnectDatabase()
	})
}

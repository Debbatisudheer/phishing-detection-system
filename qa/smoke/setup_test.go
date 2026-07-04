package smoke

import (
	"sync"

	"phishing-platform/database"
)

var once sync.Once

func setupDatabase() {

	once.Do(func() {

		database.ConnectDatabase()

	})

}
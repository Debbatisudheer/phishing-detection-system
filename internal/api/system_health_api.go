package api

import (
	"encoding/json"
	"net/http"

	"phishing-platform/database"
)

func GetSystemHealthHandler(
	w http.ResponseWriter,
	r *http.Request,
) {

	lastCleanup := "Never"

	if !database.LastCleanupTime.IsZero() {

		lastCleanup =
			database.LastCleanupTime.Format(
				"03:04:05 PM",
			)

	}

	response := map[string]interface{}{

		"database_status": "Healthy",

		"auto_cleanup": "Enabled",

		"retention": "30 Minutes",

		"last_cleanup": lastCleanup,
	}

	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	json.NewEncoder(w).Encode(
		response,
	)
}
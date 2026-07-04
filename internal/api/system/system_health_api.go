package system

import (
	"encoding/json"
	"net/http"

	"phishing-platform/database"
)

func GetSystemHealthHandler(
	w http.ResponseWriter,
	r *http.Request,
) {

	// Only GET allowed
	if r.Method != http.MethodGet {

		http.Error(
			w,
			"Method Not Allowed",
			http.StatusMethodNotAllowed,
		)

		return
	}

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

	w.WriteHeader(
		http.StatusOK,
	)

	err := json.NewEncoder(
		w,
	).Encode(
		response,
	)

	if err != nil {

		http.Error(
			w,
			"Failed to encode response",
			http.StatusInternalServerError,
		)

		return
	}
}
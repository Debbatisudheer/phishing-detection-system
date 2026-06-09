package api

import (
	"encoding/json"
	"net/http"

	"phishing-platform/database"
)

func AlertsHandler(
	w http.ResponseWriter,
	r *http.Request,
) {

	alerts, err :=
		database.GetAlerts()

	if err != nil {

		http.Error(
			w,
			err.Error(),
			http.StatusInternalServerError,
		)

		return
	}

	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	json.NewEncoder(w).Encode(
		alerts,
	)
}
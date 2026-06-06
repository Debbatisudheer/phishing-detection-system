package api

import (
	"encoding/json"
	"net/http"

	"phishing-platform/database"
)

func RecentIncidentsHandler(
	w http.ResponseWriter,
	r *http.Request,
) {

	incidents, err :=
		database.GetRecentIncidents()

	if err != nil {

		http.Error(
			w,
			err.Error(),
			http.StatusInternalServerError,
		)

		return
	}

	json.NewEncoder(w).Encode(
		incidents,
	)
}
package api

import (
	"encoding/json"
	"net/http"

	"phishing-platform/database"
)

func MITREStatsHandler(
	w http.ResponseWriter,
	r *http.Request,
) {

	stats, err :=
		database.GetMITREStats()

	if err != nil {

		http.Error(
			w,
			err.Error(),
			http.StatusInternalServerError,
		)

		return
	}

	json.NewEncoder(w).Encode(
		stats,
	)
}
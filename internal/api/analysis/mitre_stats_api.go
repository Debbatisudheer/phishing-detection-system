package api

import (
	"encoding/json"
	"net/http"

	mitrerepo "phishing-platform/database/mitre"
)

func MITREStatsHandler(
	w http.ResponseWriter,
	r *http.Request,
) {

	stats, err :=
		mitrerepo.GetMITREStats()

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
package api

import (
	"encoding/json"
	"net/http"

	threatrepo "phishing-platform/database/threathunting"
)

func ThreatHuntingHandler(
	w http.ResponseWriter,
	r *http.Request,
) {

	data, err :=
		threatrepo.GetThreatHuntingStats()

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
		data,
	)
}
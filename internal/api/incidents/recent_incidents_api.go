package api

import (
	"encoding/json"
	"net/http"

	 incidentrepo "phishing-platform/database/incidents"
)

func RecentIncidentsHandler(
	w http.ResponseWriter,
	r *http.Request,
) {

	incidents, err :=
		incidentrepo.GetRecentIncidents()

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
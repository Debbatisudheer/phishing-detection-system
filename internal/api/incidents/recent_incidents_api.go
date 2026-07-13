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

	// Only GET allowed
	if r.Method != http.MethodGet {

		http.Error(
			w,
			"Method NotAllowed",
			http.StatusMethodNotAllowed,
		)

		return
	}

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

	if incidents == nil {

		incidents = []map[string]interface{}{}

	}

	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	w.WriteHeader(
		http.StatusOK,
	)

	json.NewEncoder(
		w,
	).Encode(
		incidents,
	)
}

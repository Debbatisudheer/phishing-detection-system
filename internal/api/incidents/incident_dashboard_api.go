package api

import (
	"encoding/json"
	"net/http"

	incidentrepo "phishing-platform/database/incidents"
)

func IncidentDashboardHandler(
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

	stats, err :=
		incidentrepo.GetIncidentStats()

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

	w.WriteHeader(
		http.StatusOK,
	)

	json.NewEncoder(
		w,
	).Encode(
		stats,
	)
}

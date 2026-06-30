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

	json.NewEncoder(w).Encode(
		stats,
	)
}
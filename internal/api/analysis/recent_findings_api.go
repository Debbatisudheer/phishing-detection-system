package api

import (
	"encoding/json"
	"net/http"

	dashboardrepo "phishing-platform/database/dashboard"
)

func RecentFindingsHandler(
	w http.ResponseWriter,
	r *http.Request,
) {

	results, err :=
		dashboardrepo.GetRecentFindings()

	if err != nil {

		http.Error(
			w,
			err.Error(),
			http.StatusInternalServerError,
		)

		return
	}

	// Never return null
	if results == nil {

		results = []map[string]interface{}{}

	}

	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	json.NewEncoder(w).Encode(
		results,
	)
}
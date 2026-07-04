package api

import (
	"encoding/json"
	"net/http"

	dashboardrepo "phishing-platform/database/dashboard"
)

type DashboardResponse struct {
	TotalAnalyzed int `json:"total_analyzed"`
	Allow         int `json:"allow"`
	Suspicious    int `json:"suspicious"`
	Quarantine    int `json:"quarantine"`
	Critical      int `json:"critical"`
}

func DashboardHandler(
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

	total,
		allow,
		suspicious,
		quarantine,
		critical,
		err :=
		dashboardrepo.GetDashboardStats()

	if err != nil {

		http.Error(
			w,
			err.Error(),
			http.StatusInternalServerError,
		)

		return
	}

	response := DashboardResponse{
		TotalAnalyzed: total,
		Allow:         allow,
		Suspicious:    suspicious,
		Quarantine:    quarantine,
		Critical:      critical,
	}

	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	w.WriteHeader(
		http.StatusOK,
	)

	json.NewEncoder(w).Encode(
		response,
	)
}
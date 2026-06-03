package api

import (
	"encoding/json"
	"net/http"

	"phishing-platform/database"
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

	total,
		allow,
		suspicious,
		quarantine,
		critical,
		err :=
		database.GetDashboardStats()

	if err != nil {

		http.Error(
			w,
			err.Error(),
			http.StatusInternalServerError,
		)

		return
	}

	response :=
		DashboardResponse{
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

	json.NewEncoder(w).Encode(
		response,
	)
}
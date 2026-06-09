package api

import (
	"encoding/json"
	"net/http"

	"phishing-platform/database"
)

func CampaignTimelineHandler(
	w http.ResponseWriter,
	r *http.Request,
) {

	data, err :=
		database.GetCampaignTimeline()

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
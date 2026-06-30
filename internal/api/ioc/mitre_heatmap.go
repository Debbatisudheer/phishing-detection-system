package api

import (
	"encoding/json"
	"net/http"

	mitrerepo "phishing-platform/database/mitre"
)

func MITREHeatmapHandler(
	w http.ResponseWriter,
	r *http.Request,
) {

	data, err :=
		mitrerepo.GetMITREHeatmap()

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
package api

import (
	"encoding/json"
	"net/http"

	"phishing-platform/database"
)

func RecentFindingsHandler(
	w http.ResponseWriter,
	r *http.Request,
) {

	results, err :=
		database.GetAllAnalysisResults()

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
		results,
	)
}
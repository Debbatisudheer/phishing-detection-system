package api

import (
	"encoding/json"
	"net/http"

	"phishing-platform/database"
)

func SearchHandler(
	w http.ResponseWriter,
	r *http.Request,
) {

	query :=
		r.URL.Query().Get(
			"q",
		)

	results, err :=
		database.SearchAnalysisResults(
			query,
		)

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
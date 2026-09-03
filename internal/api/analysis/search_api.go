package api

import (
	"encoding/json"
	"net/http"

	searchrepo "phishing-platform/database/search"
)

func SearchHandler(
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

	query := r.URL.Query().Get("q")

	if query == "" {

		http.Error(
			w,
			"search query is required",
			http.StatusBadRequest,
		)

		return
	}

	results, err :=
		searchrepo.SearchAnalysisResults(
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

	w.WriteHeader(
		http.StatusOK,
	)

	json.NewEncoder(w).Encode(
		results,
	)
}

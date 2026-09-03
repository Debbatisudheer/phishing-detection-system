package api

import (
	"encoding/json"
	"net/http"

	analysisrepo "phishing-platform/database/analysis"
)

func HighRiskFilesHandler(
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

	results, err :=
		analysisrepo.GetHighRiskAnalysisResults()

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

	w.WriteHeader(
		http.StatusOK,
	)

	json.NewEncoder(w).Encode(
		results,
	)
}

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

	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	json.NewEncoder(w).Encode(
		results,
	)
}
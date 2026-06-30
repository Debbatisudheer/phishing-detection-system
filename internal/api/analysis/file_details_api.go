package api

import (
	"encoding/json"
	"net/http"
	"strings"
	analysisrepo "phishing-platform/database/analysis"
)

func FileDetailsHandler(
	w http.ResponseWriter,
	r *http.Request,
) {

	fileName :=
		strings.TrimPrefix(
			r.URL.Path,
			"/api/file/",
		)

	result, err :=
		analysisrepo.GetAnalysisResultByFileName(
			fileName,
		)

	if err != nil {

		http.Error(
			w,
			"File not found",
			http.StatusNotFound,
		)

		return
	}

	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	json.NewEncoder(w).Encode(
		result,
	)
}
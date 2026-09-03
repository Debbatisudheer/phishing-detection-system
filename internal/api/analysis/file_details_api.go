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

	// Only GET allowed
	if r.Method != http.MethodGet {

		http.Error(
			w,
			"Method Not Allowed",
			http.StatusMethodNotAllowed,
		)

		return
	}

	fileName :=
		strings.TrimPrefix(
			r.URL.Path,
			"/api/file/",
		)

	if fileName == "" {

		http.Error(
			w,
			"file name is required",
			http.StatusBadRequest,
		)

		return
	}

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

	w.WriteHeader(
		http.StatusOK,
	)

	json.NewEncoder(w).Encode(
		result,
	)
}

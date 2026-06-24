package handlers

import (
	"encoding/json"
	"net/http"

	"phishing-platform/database"
)

func GetSandboxReportsHandler(
	w http.ResponseWriter,
	r *http.Request,
) {

	reports, err :=
		database.GetSandboxReports()

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
		reports,
	)
}
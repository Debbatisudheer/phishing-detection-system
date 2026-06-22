package api

import (
	"encoding/json"
	"net/http"

	"phishing-platform/database"
)

func GetSandboxJobsHandler(
	w http.ResponseWriter,
	r *http.Request,
) {

	jobs, err :=
		database.GetSandboxJobs()

	if err != nil {

		http.Error(
			w,
			err.Error(),
			http.StatusInternalServerError,
		)

		return
	}

	json.NewEncoder(w).Encode(
		jobs,
	)
}
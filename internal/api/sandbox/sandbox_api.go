package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	sandboxrepo "phishing-platform/database/sandbox"
)

func GetSandboxJobsHandler(
	w http.ResponseWriter,
	r *http.Request,
) {

	jobs, err :=
		sandboxrepo.GetSandboxJobs()

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

func GetSandboxReportHandler(
	w http.ResponseWriter,
	r *http.Request,
) {

	jobIDString := strings.TrimPrefix(
		r.URL.Path,
		"/api/sandbox-report/",
	)

	jobID, err := strconv.Atoi(
		jobIDString,
	)

	if err != nil {

		http.Error(
			w,
			"Invalid Job ID",
			http.StatusBadRequest,
		)

		return
	}

	report, err :=
		sandboxrepo.GetSandboxReportByJobID(
			jobID,
		)

	if err != nil {

		http.Error(
			w,
			"Sandbox Report Not Ready",
			http.StatusNotFound,
		)

		return
	}

	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	json.NewEncoder(w).Encode(
		report,
	)
}
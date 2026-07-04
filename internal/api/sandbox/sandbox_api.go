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

	// Only GET allowed
	if r.Method != http.MethodGet {

		http.Error(
			w,
			"Method Not Allowed",
			http.StatusMethodNotAllowed,
		)

		return
	}

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

	

	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	w.WriteHeader(
		http.StatusOK,
	)

	json.NewEncoder(
		w,
	).Encode(
		jobs,
	)
}

func GetSandboxReportHandler(
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

	jobIDString := strings.TrimPrefix(
		r.URL.Path,
		"/api/sandbox-report/",
	)

	if jobIDString == "" {

		http.Error(
			w,
			"Missing Job ID",
			http.StatusBadRequest,
		)

		return
	}

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

	w.WriteHeader(
		http.StatusOK,
	)

	json.NewEncoder(
		w,
	).Encode(
		report,
	)
}
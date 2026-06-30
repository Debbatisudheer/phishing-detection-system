package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	sandboxrepo "phishing-platform/database/sandbox"
)

func GetSandboxReportHandler(
	w http.ResponseWriter,
	r *http.Request,
) {

	idStr :=
		strings.TrimPrefix(
			r.URL.Path,
			"/api/sandbox/report/",
		)

	id, err :=
		strconv.Atoi(idStr)

	if err != nil {

		http.Error(
			w,
			"Invalid Report ID",
			http.StatusBadRequest,
		)

		return
	}

	report, err :=
		sandboxrepo.GetSandboxReportByID(
			id,
		)

	if err != nil {

		http.Error(
			w,
			err.Error(),
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
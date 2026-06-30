package api

import (
	"encoding/json"
	"net/http"

	exportrepo "phishing-platform/database/export"
)

func ExportIOCsHandler(
	w http.ResponseWriter,
	r *http.Request,
) {

	results, err :=
		exportrepo.ExportIOCs()

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
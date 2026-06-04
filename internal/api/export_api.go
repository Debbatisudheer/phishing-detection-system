package api

import (
	"encoding/json"
	"net/http"

	"phishing-platform/database"
)

func ExportIOCsHandler(
	w http.ResponseWriter,
	r *http.Request,
) {

	results, err :=
		database.ExportIOCs()

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
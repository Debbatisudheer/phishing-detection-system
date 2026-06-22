package api

import (
	"encoding/json"
	"net/http"

	"phishing-platform/database"
)

func CorrelationHandler(
	w http.ResponseWriter,
	r *http.Request,
) {

	data, err :=
		database.GetCorrelatedIOCsDetailed()

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

	err = json.NewEncoder(w).Encode(
		data,
	)

	if err != nil {

		http.Error(
			w,
			err.Error(),
			http.StatusInternalServerError,
		)

		return
	}
}
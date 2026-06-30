package api

import (
	"encoding/json"
	"net/http"
	investigationrepo "phishing-platform/database/investigation"
)

func CorrelationHandler(
	w http.ResponseWriter,
	r *http.Request,
) {

	data, err :=
		investigationrepo.GetCorrelatedIOCsDetailed()

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
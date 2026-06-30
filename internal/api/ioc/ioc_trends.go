package api

import (
    "encoding/json"
    "net/http"

    threatrepo "phishing-platform/database/threatintel"
)

func IOCTrendsHandler(
	w http.ResponseWriter,
	r *http.Request,
) {

	data, err :=
		threatrepo.GetIOCTrends()

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
		data,
	)
}
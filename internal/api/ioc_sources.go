package api

import (
    "encoding/json"
    "net/http"

    threatrepo "phishing-platform/database/threatintel"
)

func IOCSourcesHandler(
	w http.ResponseWriter,
	r *http.Request,
) {

	ioc :=
		r.URL.Query().Get(
			"ioc",
		)

	data, err :=
		threatrepo.GetIOCSources(
			ioc,
		)

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
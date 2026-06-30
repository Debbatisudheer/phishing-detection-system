package api

import (
    "encoding/json"
    "net/http"

    alertrepo "phishing-platform/database/alerts"
)

func AlertsHandler(
	w http.ResponseWriter,
	r *http.Request,
) {

	alerts, err :=
		alertrepo.GetAlerts()

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
		alerts,
	)
}
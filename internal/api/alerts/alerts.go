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

	// Only GET allowed
	if r.Method != http.MethodGet {

		http.Error(
			w,
			"Method Not Allowed",
			http.StatusMethodNotAllowed,
		)

		return
	}

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

	if alerts == nil {

		alerts = []map[string]interface{}{}

	}

	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	w.WriteHeader(
		http.StatusOK,
	)

	err = json.NewEncoder(
		w,
	).Encode(
		alerts,
	)

	if err != nil {

		http.Error(
			w,
			"Failed to encode response",
			http.StatusInternalServerError,
		)

		return
	}
}
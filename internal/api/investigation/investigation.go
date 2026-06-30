package api

import (
	"encoding/json"
	"net/http"

	"phishing-platform/database"
)

func InvestigationSummaryHandler(
	w http.ResponseWriter,
	r *http.Request,
) {

	ioc :=
		r.URL.Query().Get(
			"ioc",
		)

	data, err :=
		database.GetInvestigationSummary(
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
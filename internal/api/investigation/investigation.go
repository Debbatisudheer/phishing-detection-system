package api

import (
	"encoding/json"
	"net/http"
	investigationrepo "phishing-platform/database/investigation"
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
		investigationrepo.GetInvestigationSummary(
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
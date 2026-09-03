package threatintel

import (
	"encoding/json"
	"net/http"

	threatrepo "phishing-platform/database/threatintel"
)

func IOCReputationHandler(
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

	ioc := r.URL.Query().Get(
		"ioc",
	)

	if ioc == "" {

		http.Error(
			w,
			"ioc is required",
			http.StatusBadRequest,
		)

		return
	}

	data, err :=
		threatrepo.GetIOCReputation(
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

	w.WriteHeader(
		http.StatusOK,
	)

	json.NewEncoder(w).Encode(
		data,
	)
}

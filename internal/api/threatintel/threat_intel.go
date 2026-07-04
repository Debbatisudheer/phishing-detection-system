package threatintel

import (
	"encoding/json"
	"net/http"

	threatrepo "phishing-platform/database/threatintel"
)

func ThreatIntelHandler(
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

	stats, err :=
		threatrepo.GetThreatIntelStats()

	if err != nil {

		http.Error(
			w,
			err.Error(),
			http.StatusInternalServerError,
		)

		return
	}

	files, err :=
		threatrepo.GetTopRiskFiles()

	if err != nil {

		http.Error(
			w,
			err.Error(),
			http.StatusInternalServerError,
		)

		return
	}

	iocs, err :=
		threatrepo.GetTopIOCs()

	if err != nil {

		http.Error(
			w,
			err.Error(),
			http.StatusInternalServerError,
		)

		return
	}

	if files == nil {
		files = []map[string]interface{}{}
	}

	if iocs == nil {
		iocs = []map[string]interface{}{}
	}

	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	w.WriteHeader(
		http.StatusOK,
	)

	json.NewEncoder(w).Encode(
		map[string]interface{}{
			"stats":     stats,
			"top_files": files,
			"top_iocs":  iocs,
		},
	)
}
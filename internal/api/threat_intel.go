package api

import (
	"encoding/json"
	"net/http"

	threatrepo "phishing-platform/database/threatintel"
)

func ThreatIntelHandler(
	w http.ResponseWriter,
	r *http.Request,
) {

	stats, _ :=
		threatrepo.GetThreatIntelStats()

	files, _ :=
		threatrepo.GetTopRiskFiles()

	iocs, _ :=
		threatrepo.GetTopIOCs()

	json.NewEncoder(w).Encode(
		map[string]interface{}{
			"stats": stats,
			"top_files": files,
			"top_iocs": iocs,
		},
	)
}
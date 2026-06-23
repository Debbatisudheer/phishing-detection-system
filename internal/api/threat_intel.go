package api

import (
	"encoding/json"
	"net/http"

	"phishing-platform/database"
)

func ThreatIntelHandler(
	w http.ResponseWriter,
	r *http.Request,
) {

	stats, _ :=
		database.GetThreatIntelStats()

	files, _ :=
		database.GetTopRiskFiles()

	iocs, _ :=
		database.GetTopIOCs()

	json.NewEncoder(w).Encode(
		map[string]interface{}{
			"stats": stats,
			"top_files": files,
			"top_iocs": iocs,
		},
	)
}
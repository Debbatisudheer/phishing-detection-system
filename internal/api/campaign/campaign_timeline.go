package campaign

import (
	"encoding/json"
	"net/http"

	campaignrepo "phishing-platform/database/campaign"
)

func CampaignTimelineHandler(
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

	data, err :=
		campaignrepo.GetCampaignTimeline()

	if err != nil {

		http.Error(
			w,
			err.Error(),
			http.StatusInternalServerError,
		)

		return
	}

	if data == nil {

		data = []map[string]interface{}{}

	}

	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	w.WriteHeader(
		http.StatusOK,
	)

	json.NewEncoder(
		w,
	).Encode(
		data,
	)
}

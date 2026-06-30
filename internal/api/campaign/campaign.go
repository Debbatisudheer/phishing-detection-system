package campaign

import (
	"encoding/json"
	"net/http"
	campaignrepo "phishing-platform/database/campaign"
)

func CampaignHandler(
	w http.ResponseWriter,
	r *http.Request,
) {

	data, err :=
		campaignrepo.GetCampaigns()

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
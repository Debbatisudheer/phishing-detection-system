package routes

import (
	"net/http"

	campaignapi "phishing-platform/internal/api/campaign"
)

func RegisterCampaignRoutes() {

	http.HandleFunc(
		"/api/campaigns",
		campaignapi.CampaignHandler,
	)

	http.HandleFunc(
		"/api/campaign-stats",
		campaignapi.CampaignStatsHandler,
	)

	http.HandleFunc(
		"/api/campaign-timeline",
		campaignapi.CampaignTimelineHandler,
	)

}
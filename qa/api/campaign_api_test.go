package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	campaignapi "phishing-platform/internal/api/campaign"
)

func TestCampaignAPI(t *testing.T) {

	setupDatabase()

	tests := []struct {
		name           string
		handler        http.HandlerFunc
		method         string
		url            string
		expectedStatus int
	}{
		{
			name:           "Campaign GET",
			handler:        campaignapi.CampaignHandler,
			method:         http.MethodGet,
			url:            "/api/campaign",
			expectedStatus: http.StatusOK,
		},
		{
			name:           "Campaign Invalid Method",
			handler:        campaignapi.CampaignHandler,
			method:         http.MethodPost,
			url:            "/api/campaign",
			expectedStatus: http.StatusMethodNotAllowed,
		},
		{
			name:           "Campaign Timeline GET",
			handler:        campaignapi.CampaignTimelineHandler,
			method:         http.MethodGet,
			url:            "/api/campaign-timeline",
			expectedStatus: http.StatusOK,
		},
		{
			name:           "Campaign Timeline Invalid Method",
			handler:        campaignapi.CampaignTimelineHandler,
			method:         http.MethodPost,
			url:            "/api/campaign-timeline",
			expectedStatus: http.StatusMethodNotAllowed,
		},
		{
			name:           "Campaign Stats GET",
			handler:        campaignapi.CampaignStatsHandler,
			method:         http.MethodGet,
			url:            "/api/campaign-stats",
			expectedStatus: http.StatusOK,
		},
		{
			name:           "Campaign Stats Invalid Method",
			handler:        campaignapi.CampaignStatsHandler,
			method:         http.MethodPost,
			url:            "/api/campaign-stats",
			expectedStatus: http.StatusMethodNotAllowed,
		},
	}

	for _, tc := range tests {

		t.Run(tc.name, func(t *testing.T) {

			req := httptest.NewRequest(
				tc.method,
				tc.url,
				nil,
			)

			rec := httptest.NewRecorder()

			tc.handler(
				rec,
				req,
			)

			if rec.Code != tc.expectedStatus {

				t.Fatalf(
					"%s expected %d got %d",
					tc.name,
					tc.expectedStatus,
					rec.Code,
				)
			}
		})
	}
}

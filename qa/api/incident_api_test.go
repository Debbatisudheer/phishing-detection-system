package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	incidentapi "phishing-platform/internal/api/incidents"
)

func TestIncidentAPI(t *testing.T) {

	setupDatabase()

	tests := []struct {
		name           string
		handler        http.HandlerFunc
		method         string
		url            string
		expectedStatus int
	}{
		{
			name:           "Incident Dashboard GET",
			handler:        incidentapi.IncidentDashboardHandler,
			method:         http.MethodGet,
			url:            "/api/incident-dashboard",
			expectedStatus: http.StatusOK,
		},
		{
			name:           "Incident Dashboard Invalid Method",
			handler:        incidentapi.IncidentDashboardHandler,
			method:         http.MethodPost,
			url:            "/api/incident-dashboard",
			expectedStatus: http.StatusMethodNotAllowed,
		},
		{
			name:           "Recent Incidents GET",
			handler:        incidentapi.RecentIncidentsHandler,
			method:         http.MethodGet,
			url:            "/api/recent-incidents",
			expectedStatus: http.StatusOK,
		},
		{
			name:           "Recent Incidents Invalid Method",
			handler:        incidentapi.RecentIncidentsHandler,
			method:         http.MethodPost,
			url:            "/api/recent-incidents",
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

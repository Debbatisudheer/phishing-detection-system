package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	alertsapi "phishing-platform/internal/api/alerts"
)

func TestAlertsAPI(t *testing.T) {

	setupDatabase()

	tests := []struct {
		name           string
		method         string
		expectedStatus int
	}{
		{
			name:           "Alerts GET",
			method:         http.MethodGet,
			expectedStatus: http.StatusOK,
		},
		{
			name:           "Alerts Invalid Method",
			method:         http.MethodPost,
			expectedStatus: http.StatusMethodNotAllowed,
		},
	}

	for _, tc := range tests {

		t.Run(tc.name, func(t *testing.T) {

			req := httptest.NewRequest(
				tc.method,
				"/api/alerts",
				nil,
			)

			rec := httptest.NewRecorder()

			alertsapi.AlertsHandler(
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

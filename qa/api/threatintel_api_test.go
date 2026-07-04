package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	threatapi "phishing-platform/internal/api/threatintel"
)

func TestThreatIntelAPI(t *testing.T) {

	setupDatabase()

	tests := []struct {
		name           string
		handler        http.HandlerFunc
		method         string
		url            string
		expectedStatus int
	}{
		{
			name:           "Threat Intel GET",
			handler:        threatapi.ThreatIntelHandler,
			method:         http.MethodGet,
			url:            "/api/threat-intel",
			expectedStatus: http.StatusOK,
		},
		{
			name:           "Threat Intel Invalid Method",
			handler:        threatapi.ThreatIntelHandler,
			method:         http.MethodPost,
			url:            "/api/threat-intel",
			expectedStatus: http.StatusMethodNotAllowed,
		},
		{
			name:           "IOC Reputation Missing IOC",
			handler:        threatapi.IOCReputationHandler,
			method:         http.MethodGet,
			url:            "/api/ioc-reputation",
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "IOC Reputation Invalid Method",
			handler:        threatapi.IOCReputationHandler,
			method:         http.MethodPost,
			url:            "/api/ioc-reputation?ioc=evil.com",
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
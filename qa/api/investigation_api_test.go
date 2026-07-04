package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	investigationapi "phishing-platform/internal/api/investigation"
)

func TestInvestigationAPI(t *testing.T) {

	setupDatabase()

	tests := []struct {
		name           string
		handler        http.HandlerFunc
		method         string
		url            string
		expectedStatus int
	}{
		{
			name:           "Investigation Missing IOC",
			handler:        investigationapi.InvestigationSummaryHandler,
			method:         http.MethodGet,
			url:            "/api/investigation",
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "Investigation Invalid Method",
			handler:        investigationapi.InvestigationSummaryHandler,
			method:         http.MethodPost,
			url:            "/api/investigation?ioc=evil.com",
			expectedStatus: http.StatusMethodNotAllowed,
		},
		{
			name:           "Correlation GET",
			handler:        investigationapi.CorrelationHandler,
			method:         http.MethodGet,
			url:            "/api/correlation",
			expectedStatus: http.StatusOK,
		},
		{
			name:           "Correlation Invalid Method",
			handler:        investigationapi.CorrelationHandler,
			method:         http.MethodPost,
			url:            "/api/correlation",
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
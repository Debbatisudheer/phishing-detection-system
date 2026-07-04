package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	analysisapi "phishing-platform/internal/api/analysis"
)

func TestExportIOCsAPI(t *testing.T) {

	setupDatabase()

	tests := []struct {
		name           string
		method         string
		expectedStatus int
	}{
		{
			name:           "Valid GET",
			method:         http.MethodGet,
			expectedStatus: http.StatusOK,
		},
		{
			name:           "Invalid Method",
			method:         http.MethodPost,
			expectedStatus: http.StatusMethodNotAllowed,
		},
	}

	for _, tc := range tests {

		t.Run(tc.name, func(t *testing.T) {

			req := httptest.NewRequest(
				tc.method,
				"/api/export/iocs",
				nil,
			)

			rec := httptest.NewRecorder()

			analysisapi.ExportIOCsHandler(
				rec,
				req,
			)

			if rec.Code != tc.expectedStatus {

				t.Fatalf(
					"expected %d got %d",
					tc.expectedStatus,
					rec.Code,
				)
			}
		})
	}
}
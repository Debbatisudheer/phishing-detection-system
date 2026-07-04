package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	reportapi "phishing-platform/internal/api/report"
)

func TestReportAPI(t *testing.T) {

	setupDatabase()

	tests := []struct {
		name           string
		method         string
		expectedStatus int
	}{
		{
			name:           "GET Report",
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
				"/api/export-report",
				nil,
			)

			rec := httptest.NewRecorder()

			reportapi.ExportReportHandler(
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
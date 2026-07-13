package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"phishing-platform/database"
	reportapi "phishing-platform/internal/api/report"
)

func TestReportAPI(t *testing.T) {

	setupDatabase()

	// Clean old data
	_, err := database.DB.Exec(`
		DELETE FROM sandbox_reports
	`)
	if err != nil {
		t.Fatal(err)
	}

	// Insert one report for testing
	_, err = database.DB.Exec(`
		INSERT INTO sandbox_reports
		(
			file_name,
			risk_score,
			risk_level,
			verdict,
			findings,
			mitre
		)
		VALUES
		(
			'sample.pdf',
			900,
			'CRITICAL',
			'QUARANTINE',
			'Test Findings',
			'T1566'
		)
	`)
	if err != nil {
		t.Fatal(err)
	}

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
					"%s expected %d got %d\nResponse: %s",
					tc.name,
					tc.expectedStatus,
					rec.Code,
					rec.Body.String(),
				)
			}
		})
	}
}

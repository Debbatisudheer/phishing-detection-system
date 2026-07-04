package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	sandboxapi "phishing-platform/internal/api/sandbox"
)

func TestSandboxAPI(t *testing.T) {

	setupDatabase()

	tests := []struct {
		name           string
		handler        http.HandlerFunc
		method         string
		url            string
		expectedStatus int
	}{
		{
			name:           "Sandbox Jobs GET",
			handler:        sandboxapi.GetSandboxJobsHandler,
			method:         http.MethodGet,
			url:            "/api/sandbox-jobs",
			expectedStatus: http.StatusOK,
		},
		{
			name:           "Sandbox Jobs Invalid Method",
			handler:        sandboxapi.GetSandboxJobsHandler,
			method:         http.MethodPost,
			url:            "/api/sandbox-jobs",
			expectedStatus: http.StatusMethodNotAllowed,
		},
		{
			name:           "Sandbox Report Missing ID",
			handler:        sandboxapi.GetSandboxReportHandler,
			method:         http.MethodGet,
			url:            "/api/sandbox-report/",
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "Sandbox Report Invalid ID",
			handler:        sandboxapi.GetSandboxReportHandler,
			method:         http.MethodGet,
			url:            "/api/sandbox-report/abc",
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "Sandbox Report Invalid Method",
			handler:        sandboxapi.GetSandboxReportHandler,
			method:         http.MethodPost,
			url:            "/api/sandbox-report/1",
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
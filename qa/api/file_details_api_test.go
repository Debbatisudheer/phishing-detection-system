package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	analysisapi "phishing-platform/internal/api/analysis"
)

func TestFileDetailsAPI(t *testing.T) {

	setupDatabase()

	tests := []struct {
		name           string
		method         string
		url            string
		expectedStatus int
	}{
		{
			name:           "Invalid Method",
			method:         http.MethodPost,
			url:            "/api/file/test.eml",
			expectedStatus: http.StatusMethodNotAllowed,
		},
		{
			name:           "Missing File Name",
			method:         http.MethodGet,
			url:            "/api/file/",
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "Unknown File",
			method:         http.MethodGet,
			url:            "/api/file/this_file_does_not_exist.eml",
			expectedStatus: http.StatusNotFound,
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

			analysisapi.FileDetailsHandler(
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
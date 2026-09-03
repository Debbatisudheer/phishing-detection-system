package api

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	emailapi "phishing-platform/internal/api/email"
)

func TestAnalyzeEmailAPI(t *testing.T) {

	setupDatabase()

	tests := []struct {
		name           string
		method         string
		body           string
		expectedStatus int
	}{
		{
			name:           "Invalid Method",
			method:         http.MethodGet,
			body:           "",
			expectedStatus: http.StatusMethodNotAllowed,
		},
		{
			name:           "Invalid JSON",
			method:         http.MethodPost,
			body:           `{`,
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "Empty Request",
			method:         http.MethodPost,
			body:           `{}`,
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "Safe Email",
			method:         http.MethodPost,
			body:           `{"subject":"Meeting","body":"Hello team"}`,
			expectedStatus: http.StatusOK,
		},
		{
			name:           "Email With URL",
			method:         http.MethodPost,
			body:           `{"subject":"Verify Account","body":"Click https://evil.com/login immediately"}`,
			expectedStatus: http.StatusOK,
		},
		{
			name:           "Email Without URL",
			method:         http.MethodPost,
			body:           `{"subject":"Greetings","body":"How are you?"}`,
			expectedStatus: http.StatusOK,
		},
	}

	for _, tc := range tests {

		t.Run(tc.name, func(t *testing.T) {

			req := httptest.NewRequest(
				tc.method,
				"/api/analyze-email",
				bytes.NewBufferString(tc.body),
			)

			rec := httptest.NewRecorder()

			emailapi.AnalyzeEmailHandler(
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

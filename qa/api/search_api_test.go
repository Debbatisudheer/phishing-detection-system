package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	analysisapi "phishing-platform/internal/api/analysis"
)

func TestSearchAPI(t *testing.T) {

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
			url:            "/api/search?q=paypal",
			expectedStatus: http.StatusMethodNotAllowed,
		},
		{
			name:           "Missing Query",
			method:         http.MethodGet,
			url:            "/api/search",
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "Unknown Search",
			method:         http.MethodGet,
			url:            "/api/search?q=this_should_not_exist_123456",
			expectedStatus: http.StatusOK,
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

			analysisapi.SearchHandler(
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
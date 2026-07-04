package api

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	authapi "phishing-platform/internal/api/auth"
)

func TestRegisterAPI(t *testing.T) {

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
			name:           "Empty Username",
			method:         http.MethodPost,
			body:           `{"username":"","password":"password123"}`,
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "Empty Password",
			method:         http.MethodPost,
			body:           `{"username":"sudheer","password":""}`,
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "Empty Body",
			method:         http.MethodPost,
			body:           `{}`,
			expectedStatus: http.StatusBadRequest,
		},
	}

	for _, tc := range tests {

		t.Run(tc.name, func(t *testing.T) {

			req := httptest.NewRequest(
				tc.method,
				"/api/register",
				bytes.NewBufferString(tc.body),
			)

			rec := httptest.NewRecorder()

			authapi.RegisterHandler(
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
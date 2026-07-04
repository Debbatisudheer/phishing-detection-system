package api

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
	caseapi "phishing-platform/internal/api/cases"
)

func TestCaseAPI(t *testing.T) {

	setupDatabase()

	tests := []struct {
		name           string
		handler        http.HandlerFunc
		method         string
		url            string
		body           string
		expectedStatus int
	}{
		{
			name:           "Create Invalid Method",
			handler:        caseapi.CreateCaseHandler,
			method:         http.MethodGet,
			url:            "/api/case",
			body:           "",
			expectedStatus: http.StatusMethodNotAllowed,
		},
		{
			name:           "Create Invalid JSON",
			handler:        caseapi.CreateCaseHandler,
			method:         http.MethodPost,
			url:            "/api/case",
			body:           `{`,
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "Create Missing Fields",
			handler:        caseapi.CreateCaseHandler,
			method:         http.MethodPost,
			url:            "/api/case",
			body:           `{}`,
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "Get Cases",
			handler:        caseapi.GetCasesHandler,
			method:         http.MethodGet,
			url:            "/api/cases",
			body:           "",
			expectedStatus: http.StatusOK,
		},
		{
			name:           "Update Invalid JSON",
			handler:        caseapi.UpdateCaseHandler,
			method:         http.MethodPut,
			url:            "/api/case/1",
			body:           `{`,
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "Unknown Case",
			handler:        caseapi.GetCaseHandler,
			method:         http.MethodGet,
			url:            "/api/case-details/999999",
			body:           "",
			expectedStatus: http.StatusNotFound,
		},
		{
			name:           "Close Missing ID",
			handler:        caseapi.CloseCaseHandler,
			method:         http.MethodPost,
			url:            "/api/case-close/",
			body:           "",
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "Add Note Invalid JSON",
			handler:        caseapi.AddCaseNoteHandler,
			method:         http.MethodPost,
			url:            "/api/case-note",
			body:           `{`,
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "Add Note Missing Fields",
			handler:        caseapi.AddCaseNoteHandler,
			method:         http.MethodPost,
			url:            "/api/case-note",
			body:           `{}`,
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "Get Notes Missing ID",
			handler:        caseapi.GetCaseNotesHandler,
			method:         http.MethodGet,
			url:            "/api/case-notes/",
			body:           "",
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "Get Notes Invalid ID",
			handler:        caseapi.GetCaseNotesHandler,
			method:         http.MethodGet,
			url:            "/api/case-notes/abc",
			body:           "",
			expectedStatus: http.StatusBadRequest,
		},
	}

	for _, tc := range tests {

		t.Run(tc.name, func(t *testing.T) {

			req := httptest.NewRequest(
				tc.method,
				tc.url,
				bytes.NewBufferString(tc.body),
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


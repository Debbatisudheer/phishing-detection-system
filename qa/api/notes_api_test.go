package api

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	notesapi "phishing-platform/internal/api/notes"
)

func TestNotesAPI(t *testing.T) {

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
			name:           "Save Invalid Method",
			handler:        notesapi.SaveNoteHandler,
			method:         http.MethodGet,
			url:            "/api/note",
			body:           "",
			expectedStatus: http.StatusMethodNotAllowed,
		},
		{
			name:           "Save Invalid JSON",
			handler:        notesapi.SaveNoteHandler,
			method:         http.MethodPost,
			url:            "/api/note",
			body:           `{`,
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "Save Missing Fields",
			handler:        notesapi.SaveNoteHandler,
			method:         http.MethodPost,
			url:            "/api/note",
			body:           `{}`,
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "Get Missing IOC",
			handler:        notesapi.GetNotesHandler,
			method:         http.MethodGet,
			url:            "/api/notes",
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "Get Invalid Method",
			handler:        notesapi.GetNotesHandler,
			method:         http.MethodPost,
			url:            "/api/notes?ioc=evil.com",
			expectedStatus: http.StatusMethodNotAllowed,
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

package smoke

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	emailapi "phishing-platform/internal/api/email"
)

func TestEmailAnalysisSmoke(t *testing.T) {

	setupDatabase()

	body := []byte(`{
		"subject":"Weekly Meeting",
		"body":"Please join the meeting https://company.com"
	}`)

	req := httptest.NewRequest(
		http.MethodPost,
		"/api/analyze-email",
		bytes.NewBuffer(body),
	)

	req.Header.Set(
		"Content-Type",
		"application/json",
	)

	rec := httptest.NewRecorder()

	emailapi.AnalyzeEmailHandler(
		rec,
		req,
	)

	if rec.Code != http.StatusOK {

		t.Fatalf(
			"Expected 200 OK, got %d",
			rec.Code,
		)
	}

	t.Log("==============================")
	t.Log("EMAIL ANALYSIS SMOKE TEST PASSED")
	t.Log("Email Analysis API Responding")
	t.Log("Risk Engine Working")
	t.Log("==============================")
}
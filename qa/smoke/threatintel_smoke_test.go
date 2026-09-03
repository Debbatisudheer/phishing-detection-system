package smoke

import (
	"net/http"
	"net/http/httptest"
	"testing"

	threatintelapi "phishing-platform/internal/api/threatintel"
)

func TestThreatIntelSmoke(t *testing.T) {

	setupDatabase()

	req := httptest.NewRequest(
		http.MethodGet,
		"/api/threat-intel",
		nil,
	)

	rec := httptest.NewRecorder()

	threatintelapi.ThreatIntelHandler(
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
	t.Log("THREAT INTELLIGENCE SMOKE TEST PASSED")
	t.Log("Threat Intelligence API Responding")
	t.Log("==============================")
}

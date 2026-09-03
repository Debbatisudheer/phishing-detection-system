package smoke

import (
	"net/http"
	"net/http/httptest"
	"testing"

	systemapi "phishing-platform/internal/api/system"
)

func TestSystemHealthSmoke(t *testing.T) {

	setupDatabase()

	req := httptest.NewRequest(
		http.MethodGet,
		"/api/system-health",
		nil,
	)

	rec := httptest.NewRecorder()

	systemapi.GetSystemHealthHandler(
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
	t.Log("SYSTEM HEALTH SMOKE TEST PASSED")
	t.Log("System Health API Responding")
	t.Log("Database Healthy")
	t.Log("================================")
}

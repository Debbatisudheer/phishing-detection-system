package smoke

import (
	"net/http"
	"net/http/httptest"
	"testing"

	dashboardapi "phishing-platform/internal/api/dashboard"
)

func TestDashboardSmoke(t *testing.T) {

	setupDatabase()

	req := httptest.NewRequest(
		http.MethodGet,
		"/api/dashboard",
		nil,
	)

	rec := httptest.NewRecorder()

	dashboardapi.DashboardHandler(
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
	t.Log("DASHBOARD SMOKE TEST PASSED")
	t.Log("Dashboard API Responding")
	t.Log("==============================")
}

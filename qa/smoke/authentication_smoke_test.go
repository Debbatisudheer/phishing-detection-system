package smoke

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	authapi "phishing-platform/internal/api/auth"
)

func TestAuthenticationSmoke(t *testing.T) {

	setupDatabase()

	body := []byte(`{}`)

	req := httptest.NewRequest(
		http.MethodPost,
		"/api/login",
		bytes.NewBuffer(body),
	)

	rec := httptest.NewRecorder()

	authapi.LoginHandler(
		rec,
		req,
	)

	// We expect either:
	// 400 -> Invalid Request
	// 401 -> User not found
	//
	// Both prove the authentication service
	// is alive and responding.

	if rec.Code != http.StatusBadRequest &&
		rec.Code != http.StatusUnauthorized {

		t.Fatalf(
			"Unexpected Status Code: %d",
			rec.Code,
		)
	}

	t.Log("==============================")
	t.Log("AUTHENTICATION SMOKE TEST PASSED")
	t.Log("Authentication API Responding")
	t.Log("==============================")
}

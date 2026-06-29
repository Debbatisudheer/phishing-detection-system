package routes

import (
	"net/http"

	authapi "phishing-platform/internal/api/auth"
)

func RegisterAuthRoutes() {

	http.HandleFunc(
		"/api/register",
		authapi.RegisterHandler,
	)

	http.HandleFunc(
		"/api/login",
		authapi.LoginHandler,
	)

}
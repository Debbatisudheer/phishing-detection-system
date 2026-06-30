package routes

import (
	"net/http"

	dashboardapi "phishing-platform/internal/api/dashboard"
)

func RegisterDashboardRoutes() {

	http.HandleFunc(
		"/api/dashboard",
		dashboardapi.DashboardHandler,
	)

}
package routes

import (
	"net/http"

	reportapi "phishing-platform/internal/api/report"
	systemapi "phishing-platform/internal/api/system"
	api "phishing-platform/internal/api"
)

func RegisterSystemRoutes() {

	http.HandleFunc(
		"/health",
		api.HealthHandler,
	)

	http.HandleFunc(
		"/api/system-health",
		systemapi.GetSystemHealthHandler,
	)

	http.HandleFunc(
		"/api/export-report",
		reportapi.ExportReportHandler,
	)

}
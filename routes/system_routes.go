package routes

import (
	"net/http"

	api "phishing-platform/internal/api"
	reportapi "phishing-platform/internal/api/report"
	systemapi "phishing-platform/internal/api/system"
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

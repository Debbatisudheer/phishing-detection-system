package routes

import (
	"net/http"

	sandboxapi "phishing-platform/internal/api/sandbox"
	"phishing-platform/internal/handlers"
)

func RegisterSandboxRoutes() {

	http.HandleFunc(
		"/api/sandbox-jobs",
		sandboxapi.GetSandboxJobsHandler,
	)

	http.HandleFunc(
		"/api/sandbox-report/",
		sandboxapi.GetSandboxReportHandler,
	)

	http.HandleFunc(
		"/api/sandbox/reports",
		handlers.GetSandboxReportsHandler,
	)

	http.HandleFunc(
		"/api/sandbox/report/",
		handlers.GetSandboxReportHandler,
	)

}
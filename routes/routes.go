package routes

import (
	"net/http"

	"phishing-platform/internal/api"
	"phishing-platform/internal/handlers"
	"phishing-platform/internal/websocket"
)

func SetupRoutes() {

	http.HandleFunc(
	"/health",
	api.HealthHandler,
)

http.HandleFunc(
	"/api/analyze-email",
	api.AnalyzeEmailHandler,
)

http.HandleFunc(
	"/api/analyze-file",
	api.AnalyzeFileHandler,
)

http.HandleFunc(
	"/api/dashboard",
	api.DashboardHandler,
)

http.HandleFunc(
	"/api/recent-findings",
	api.RecentFindingsHandler,
)
http.HandleFunc(
	"/api/high-risk-files",
	api.HighRiskFilesHandler,
)
	// Home Route
	http.HandleFunc(
		"/",
		handlers.HomeHandler,
	)

	// Email Analysis API
	http.HandleFunc(
		"/analyze",
		handlers.EmailHandler,
	)

	// Fetch Emails API
	http.HandleFunc(
		"/emails",
		handlers.GetEmailsHandler,
	)

	// Quarantine API
	http.HandleFunc(
		"/quarantine/",
		handlers.QuarantineHandler,
	)

	// WebSocket Route
	http.HandleFunc(
		"/ws",
		websocket.HandleConnections,
	)

	http.HandleFunc(
	"/add-note/",
	handlers.AddNoteHandler,
)
}
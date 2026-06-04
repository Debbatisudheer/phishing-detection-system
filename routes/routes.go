package routes

import (
	"net/http"

	"phishing-platform/internal/api"
	"phishing-platform/internal/handlers"
	"phishing-platform/internal/websocket"
	"phishing-platform/internal/middleware"
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

http.HandleFunc(
	"/api/file/",
	api.FileDetailsHandler,
)

http.HandleFunc(
	"/api/search",
	middleware.JWTMiddleware(
		api.SearchHandler,
	),
)

http.HandleFunc(
	"/api/case",
	api.CreateCaseHandler,
)

http.HandleFunc(
	"/api/cases",
	middleware.JWTMiddleware(
		api.GetCasesHandler,
	),
)

http.HandleFunc(
	"/api/case/",
	api.UpdateCaseHandler,
)

http.HandleFunc(
	"/api/case-details/",
	api.GetCaseHandler,
)

http.HandleFunc(
	"/api/case-close/",
	middleware.JWTMiddleware(
		api.CloseCaseHandler,
	),
)

http.HandleFunc(
	"/api/export/iocs",
	middleware.JWTMiddleware(
		api.ExportIOCsHandler,
	),
)

http.HandleFunc(
	"/api/register",
	api.RegisterHandler,
)

http.HandleFunc(
	"/api/login",
	api.LoginHandler,
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
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
	"/api/alerts",
	api.AlertsHandler,
)

http.HandleFunc(
	"/api/recent-findings",
	middleware.JWTMiddleware(
		api.RecentFindingsHandler,
	),
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

http.HandleFunc(
	"/api/threat-hunting",
	middleware.JWTMiddleware(
		api.ThreatHuntingHandler,
	),
)

http.HandleFunc(
	"/api/mitre-stats",
	api.MITREStatsHandler,
)

http.HandleFunc(
	"/api/case-note",
	api.AddCaseNoteHandler,
)

http.HandleFunc(
	"/api/case-notes/",
	api.GetCaseNotesHandler,
)

http.HandleFunc(
	"/api/incident-dashboard",
	api.IncidentDashboardHandler,
)

http.HandleFunc(
	"/api/recent-incidents",
	api.RecentIncidentsHandler,
)

http.HandleFunc(
	"/api/correlation",
	api.CorrelationHandler,
)

http.HandleFunc(
	"/api/campaigns",
	api.CampaignHandler,
)
http.HandleFunc(
	"/api/ioc-sources",
	api.IOCSourcesHandler,
)

http.HandleFunc(
	"/api/campaign-stats",
	api.CampaignStatsHandler,
)

http.HandleFunc(
	"/api/ioc-graph",
	api.IOCGraphHandler,
)

http.HandleFunc(
	"/api/mitre-heatmap",
	api.MITREHeatmapHandler,
)

http.HandleFunc(
	"/api/campaign-timeline",
	api.CampaignTimelineHandler,
)

http.HandleFunc(
	"/api/ioc-trends",
	api.IOCTrendsHandler,
)

http.HandleFunc(
	"/api/investigation-summary",
	api.InvestigationSummaryHandler,
)

http.HandleFunc(
	"/api/notes",
	api.SaveNoteHandler,
)

http.HandleFunc(
	"/api/get-notes",
	api.GetNotesHandler,
)
http.HandleFunc(
	"/api/ioc-reputation",
	api.IOCReputationHandler,
)

http.HandleFunc(
	"/api/sandbox-jobs",
	api.GetSandboxJobsHandler,
)

http.HandleFunc(
	"/api/threat-intel",
	api.ThreatIntelHandler,
)

http.HandleFunc(
	"/api/export-report",
	api.ExportReportHandler,
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
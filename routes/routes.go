package routes

import (
    
    dashboardapi "phishing-platform/internal/api/dashboard"
    analysisapi "phishing-platform/internal/api/analysis"
    caseapi "phishing-platform/internal/api/cases"
	incidentapi "phishing-platform/internal/api/incidents"

    "phishing-platform/internal/handlers"
    "phishing-platform/internal/middleware"
    "phishing-platform/internal/websocket"
	"net/http"
	 "phishing-platform/internal/api"
)

func SetupRoutes() {

	RegisterAuthRoutes()

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
	dashboardapi.DashboardHandler,
)

http.HandleFunc(
	"/api/alerts",
	api.AlertsHandler,
)

http.HandleFunc(
	"/api/recent-findings",
	middleware.JWTMiddleware(
		analysisapi.RecentFindingsHandler,
	),
)

http.HandleFunc(
	"/api/high-risk-files",
	analysisapi.HighRiskFilesHandler,
)

http.HandleFunc(
	"/api/file/",
	analysisapi.FileDetailsHandler,
)

http.HandleFunc(
	"/api/search",
	middleware.JWTMiddleware(
		analysisapi.SearchHandler,
	),
)

http.HandleFunc(
	"/api/case",
	caseapi.CreateCaseHandler,
)

http.HandleFunc(
	"/api/cases",
	middleware.JWTMiddleware(
		caseapi.GetCasesHandler,
	),
)

http.HandleFunc(
	"/api/case/",
	caseapi.UpdateCaseHandler,
)

http.HandleFunc(
	"/api/case-details/",
	caseapi.GetCaseHandler,
)

http.HandleFunc(
	"/api/case-close/",
	middleware.JWTMiddleware(
		caseapi.CloseCaseHandler,
	),
)

http.HandleFunc(
	"/api/export/iocs",
	middleware.JWTMiddleware(
		analysisapi.ExportIOCsHandler,
	),
)


http.HandleFunc(
	"/api/threat-hunting",
	middleware.JWTMiddleware(
		analysisapi.ThreatHuntingHandler,
	),
)

http.HandleFunc(
	"/api/mitre-stats",
	analysisapi.MITREStatsHandler,
)

http.HandleFunc(
	"/api/case-note",
	caseapi.AddCaseNoteHandler,
)

http.HandleFunc(
	"/api/case-notes/",
	caseapi.GetCaseNotesHandler,
)

http.HandleFunc(
	"/api/incident-dashboard",
	incidentapi.IncidentDashboardHandler,
)

http.HandleFunc(
	"/api/recent-incidents",
	incidentapi.RecentIncidentsHandler,
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
	"/api/sandbox-report/",
	api.GetSandboxReportHandler,
)

http.HandleFunc(
	"/api/threat-intel",
	api.ThreatIntelHandler,
)

http.HandleFunc(
	"/api/export-report",
	api.ExportReportHandler,
)

http.HandleFunc(
	"/api/sandbox/reports",
	handlers.GetSandboxReportsHandler,
)
http.HandleFunc(
	"/api/sandbox/report/",
	handlers.GetSandboxReportHandler,
)

http.HandleFunc(
	"/api/system-health",
	api.GetSystemHealthHandler,
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
package routes

import (
    
    dashboardapi "phishing-platform/internal/api/dashboard"
    analysisapi "phishing-platform/internal/api/analysis"
    caseapi "phishing-platform/internal/api/cases"
	incidentapi "phishing-platform/internal/api/incidents"
	emailapi "phishing-platform/internal/api/email"
	alertsapi "phishing-platform/internal/api/alerts"
	investigationapi "phishing-platform/internal/api/investigation"
	iocapi "phishing-platform/internal/api/ioc"
	notesapi "phishing-platform/internal/api/notes"
	sandboxapi "phishing-platform/internal/api/sandbox"
	campaignapi "phishing-platform/internal/api/campaign"
	threatintelapi "phishing-platform/internal/api/threatintel"
	reportapi "phishing-platform/internal/api/report"
	systemapi "phishing-platform/internal/api/system"
	

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
	emailapi.AnalyzeEmailHandler,
)

http.HandleFunc(
	"/api/analyze-file",
	emailapi.AnalyzeFileHandler,
)

http.HandleFunc(
	"/api/dashboard",
	dashboardapi.DashboardHandler,
)

http.HandleFunc(
	"/api/alerts",
	alertsapi.AlertsHandler,
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
	investigationapi.CorrelationHandler,
)

http.HandleFunc(
	"/api/campaigns",
	campaignapi.CampaignHandler,
)
http.HandleFunc(
	"/api/ioc-sources",
	iocapi.IOCSourcesHandler,
)

http.HandleFunc(
	"/api/campaign-stats",
	campaignapi.CampaignStatsHandler,
)

http.HandleFunc(
	"/api/ioc-graph",
	iocapi.IOCGraphHandler,
)

http.HandleFunc(
	"/api/mitre-heatmap",
	iocapi.MITREHeatmapHandler,
)

http.HandleFunc(
	"/api/campaign-timeline",
	campaignapi.CampaignTimelineHandler,
)

http.HandleFunc(
	"/api/ioc-trends",
	iocapi.IOCTrendsHandler,
)

http.HandleFunc(
	"/api/investigation-summary",
	investigationapi.InvestigationSummaryHandler,
)

http.HandleFunc(
	"/api/notes",
	notesapi.SaveNoteHandler,
)

http.HandleFunc(
	"/api/get-notes",
	notesapi.GetNotesHandler,
)
http.HandleFunc(
	"/api/ioc-reputation",
	threatintelapi.IOCReputationHandler,
)

http.HandleFunc(
	"/api/sandbox-jobs",
	sandboxapi.GetSandboxJobsHandler,
)

http.HandleFunc(
	"/api/sandbox-report/",
	sandboxapi.GetSandboxReportHandler,
)

http.HandleFunc(
	"/api/threat-intel",
	threatintelapi.ThreatIntelHandler,
)

http.HandleFunc(
	"/api/export-report",
	reportapi.ExportReportHandler,
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
	systemapi.GetSystemHealthHandler,
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
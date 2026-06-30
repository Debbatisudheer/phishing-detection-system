package routes

import (
	"net/http"

	analysisapi "phishing-platform/internal/api/analysis"
	"phishing-platform/internal/middleware"
)

func RegisterAnalysisRoutes() {

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

}
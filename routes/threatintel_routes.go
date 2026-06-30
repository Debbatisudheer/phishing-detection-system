package routes

import (
	"net/http"

	alertsapi "phishing-platform/internal/api/alerts"
	investigationapi "phishing-platform/internal/api/investigation"
	iocapi "phishing-platform/internal/api/ioc"
	notesapi "phishing-platform/internal/api/notes"
	threatintelapi "phishing-platform/internal/api/threatintel"
)

func RegisterThreatIntelRoutes() {

	http.HandleFunc(
		"/api/alerts",
		alertsapi.AlertsHandler,
	)

	http.HandleFunc(
		"/api/correlation",
		investigationapi.CorrelationHandler,
	)

	http.HandleFunc(
		"/api/investigation-summary",
		investigationapi.InvestigationSummaryHandler,
	)

	http.HandleFunc(
		"/api/ioc-sources",
		iocapi.IOCSourcesHandler,
	)

	http.HandleFunc(
		"/api/ioc-graph",
		iocapi.IOCGraphHandler,
	)

	http.HandleFunc(
		"/api/ioc-trends",
		iocapi.IOCTrendsHandler,
	)

	http.HandleFunc(
		"/api/mitre-heatmap",
		iocapi.MITREHeatmapHandler,
	)

	http.HandleFunc(
		"/api/ioc-reputation",
		threatintelapi.IOCReputationHandler,
	)

	http.HandleFunc(
		"/api/threat-intel",
		threatintelapi.ThreatIntelHandler,
	)

	http.HandleFunc(
		"/api/notes",
		notesapi.SaveNoteHandler,
	)

	http.HandleFunc(
		"/api/get-notes",
		notesapi.GetNotesHandler,
	)

}
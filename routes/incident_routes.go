package routes

import (
	"net/http"

	incidentapi "phishing-platform/internal/api/incidents"
)

func RegisterIncidentRoutes() {

	http.HandleFunc(
		"/api/incident-dashboard",
		incidentapi.IncidentDashboardHandler,
	)

	http.HandleFunc(
		"/api/recent-incidents",
		incidentapi.RecentIncidentsHandler,
	)

}
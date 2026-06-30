package routes

import (
	"net/http"

	emailapi "phishing-platform/internal/api/email"
	"phishing-platform/internal/handlers"
)

func RegisterEmailRoutes() {

	http.HandleFunc(
		"/api/analyze-email",
		emailapi.AnalyzeEmailHandler,
	)

	http.HandleFunc(
		"/api/analyze-file",
		emailapi.AnalyzeFileHandler,
	)

	http.HandleFunc(
		"/",
		handlers.HomeHandler,
	)

	http.HandleFunc(
		"/analyze",
		handlers.EmailHandler,
	)

	http.HandleFunc(
		"/emails",
		handlers.GetEmailsHandler,
	)

	http.HandleFunc(
		"/quarantine/",
		handlers.QuarantineHandler,
	)

	http.HandleFunc(
		"/add-note/",
		handlers.AddNoteHandler,
	)

}
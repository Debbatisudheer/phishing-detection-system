package routes

import (
	"net/http"

	caseapi "phishing-platform/internal/api/cases"
	"phishing-platform/internal/middleware"
)

func RegisterCaseRoutes() {

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
		"/api/case-note",
		caseapi.AddCaseNoteHandler,
	)

	http.HandleFunc(
		"/api/case-notes/",
		caseapi.GetCaseNotesHandler,
	)

}
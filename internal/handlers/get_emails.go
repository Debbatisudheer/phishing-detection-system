package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"phishing-platform/database"
)

type EmailRecord struct {
	ID              int
	Sender          string
	Subject         string
	Body            string
	RiskScore       int
	Decision        string
	Findings        string
	Attachments     string
	AnalystNote     sql.NullString
	MitreTechnique  sql.NullString
}

func GetEmailsHandler(
	w http.ResponseWriter,
	r *http.Request,
) {

	rows, err := database.DB.Query(`
SELECT
	id,
	sender,
	subject,
	body,
	risk_score,
	decision,
	findings,
	attachments,
	analyst_note,
	mitre_technique
FROM public.emails
ORDER BY id DESC
`)

	if err != nil {

		http.Error(
			w,
			err.Error(),
			http.StatusInternalServerError,
		)

		return
	}

	defer rows.Close()

	var emails []map[string]interface{}

	for rows.Next() {

		var email EmailRecord

		err := rows.Scan(
			&email.ID,
			&email.Sender,
			&email.Subject,
			&email.Body,
			&email.RiskScore,
			&email.Decision,
			&email.Findings,
			&email.Attachments,
			&email.AnalystNote,
			&email.MitreTechnique,
		)

		if err != nil {

			http.Error(
				w,
				err.Error(),
				http.StatusInternalServerError,
			)

			return
		}

		// Handle NULL analyst_note
		note := ""

		if email.AnalystNote.Valid {

			note =
				email.AnalystNote.String
		}

		// Handle NULL mitre_technique
		mitreTechnique := ""

		if email.MitreTechnique.Valid {

			mitreTechnique =
				email.MitreTechnique.String
		}

		response :=
			map[string]interface{}{
				"id":               email.ID,
				"sender":           email.Sender,
				"subject":          email.Subject,
				"body":             email.Body,
				"risk_score":       email.RiskScore,
				"decision":         email.Decision,
				"findings":         email.Findings,
				"attachments":      email.Attachments,
				"analyst_note":     note,
				"mitre_technique":  mitreTechnique,
			}

		emails = append(
			emails,
			response,
		)
	}

	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	json.NewEncoder(w).Encode(emails)
}
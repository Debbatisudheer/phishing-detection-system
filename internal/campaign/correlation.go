package campaign

import (
	"phishing-platform/database"
)

func DetectCampaign(
	subject string,
) []string {

	var findings []string

	var count int

	query := `
SELECT COUNT(*)
FROM public.emails
WHERE subject = $1
`

	err := database.DB.QueryRow(
		query,
		subject,
	).Scan(&count)

	if err != nil {
		return findings
	}

	if count >= 5 {

		findings = append(
			findings,
			"Potential phishing campaign detected",
		)
	}

	return findings
}
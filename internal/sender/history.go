package sender

import (
	"phishing-platform/database"
)

func CheckSenderHistory(
	sender string,
) []string {

	var findings []string

	var count int

	query := `
SELECT COUNT(*)
FROM public.emails
WHERE sender = $1
AND decision = 'QUARANTINE'
`

	err := database.DB.QueryRow(
		query,
		sender,
	).Scan(&count)

	if err != nil {
		return findings
	}

	if count >= 3 {

		findings = append(
			findings,
			"Repeated malicious sender detected",
		)
	}

	return findings
}
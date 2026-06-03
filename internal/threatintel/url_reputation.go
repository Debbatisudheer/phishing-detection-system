package threatintel

import (
	"net/url"

	"phishing-platform/database"
)

func CheckURLReputation(
	rawURL string,
) []string {

	var findings []string

	parsedURL, err := url.Parse(rawURL)

	if err != nil {
		return findings
	}

	host := parsedURL.Hostname()

	var reputation string

	query := `
SELECT reputation
FROM url_reputation
WHERE domain = $1
`

	err = database.DB.QueryRow(
		query,
		host,
	).Scan(&reputation)

	if err != nil {
		return findings
	}

	if reputation == "malicious" {

		findings = append(
			findings,
			"URL found in reputation database",
		)
	}

	return findings
}
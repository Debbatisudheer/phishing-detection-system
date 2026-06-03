package emailauth

import "strings"

func CheckEmailAuthentication(
	sender string,
) []string {

	var findings []string

	sender = strings.ToLower(sender)

	// Simulation for now

	if strings.Contains(
		sender,
		"gmail.com",
	) {

		findings = append(
			findings,
			"SPF PASS",
		)

		findings = append(
			findings,
			"DKIM PASS",
		)

		findings = append(
			findings,
			"DMARC PASS",
		)

	} else {

		findings = append(
			findings,
			"SPF FAIL",
		)

		findings = append(
			findings,
			"DKIM FAIL",
		)

		findings = append(
			findings,
			"DMARC FAIL",
		)
	}

	return findings
}
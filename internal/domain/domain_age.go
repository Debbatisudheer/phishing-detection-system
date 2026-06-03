package domain

import "strings"

func CheckDomainAge(
	host string,
) []string {

	var findings []string

	if strings.Contains(
		host,
		"security",
	) {

		findings = append(
			findings,
			"Newly registered domain detected",
		)
	}

	return findings
}
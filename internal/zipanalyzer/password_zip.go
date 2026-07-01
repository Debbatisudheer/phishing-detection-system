package zipanalyzer

import (
	"strings"
)

func DetectPasswordProtectedZIP(
	subject string,
	body string,
) []string {

	findings := []string{}

	content :=
		strings.ToLower(
			subject + " " + body,
		)

	if strings.Contains(
		content,
		"password",
	) &&
		strings.Contains(
			content,
			"zip",
		) {

		findings = append(
			findings,
			"Password-protected ZIP suspected",
		)
	}

	return findings
}
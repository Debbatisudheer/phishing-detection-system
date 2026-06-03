package mitre

import "strings"

func MapTechnique(
	subject string,
	body string,
) string {

	subject =
		strings.ToLower(subject)

	body =
		strings.ToLower(body)

	if strings.Contains(
		body,
		"http",
	) {

		return "T1566 - Phishing"
	}

	if strings.Contains(
		body,
		"login",
	) {

		return "T1204 - User Execution"
	}

	return "T1598 - Phishing for Information"
}
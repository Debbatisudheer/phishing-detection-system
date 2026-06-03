package domain

import "strings"

func CheckRedirectURL(
	rawURL string,
) []string {

	var findings []string

	redirectServices := []string{
		"bit.ly",
		"tinyurl.com",
		"t.co",
		"goo.gl",
		"shorturl.at",
	}

	for _, service := range redirectServices {

		if strings.Contains(
			strings.ToLower(rawURL),
			service,
		) {

			findings = append(
				findings,
				"Suspicious URL redirection detected",
			)

			break
		}
	}

	return findings
}
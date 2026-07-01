package thread

import (
	"strings"
)

func DetectThreadHijack(
	subject string,
	body string,
) []string {

	findings := []string{}

	subject = strings.ToLower(
		subject,
	)

	body = strings.ToLower(
		body,
	)

	isReply :=
		strings.HasPrefix(
			subject,
			"re:",
		)

	isForward :=
		strings.HasPrefix(
			subject,
			"fw:",
		) ||
			strings.HasPrefix(
				subject,
				"fwd:",
			)

	if isReply || isForward {

		findings = append(
			findings,
			"Potential thread hijacking detected",
		)
	}

	keywords := []string{
		"change bank account",
		"new payment details",
		"wire transfer",
		"urgent payment",
		"invoice attached",
	}

	for _, keyword := range keywords {

		if strings.Contains(
			body,
			keyword,
		) {

			findings = append(
				findings,
				"Thread hijack indicator: "+keyword,
			)
		}
	}

	return findings
}
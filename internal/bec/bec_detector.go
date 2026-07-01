package bec

import (
	"strings"
)

func DetectBEC(
	subject string,
	body string,
) []string {

	findings := []string{}

	text :=
		strings.ToLower(
			subject + " " + body,
		)

	keywords := []string{
		"wire transfer",
		"bank transfer",
		"urgent payment",
		"gift card",
		"invoice payment",
		"send funds",
		"payment today",
	}

	for _, keyword := range keywords {

		if strings.Contains(
			text,
			keyword,
		) {

			findings = append(
				findings,
				"BEC indicator detected: "+keyword,
			)
		}
	}

	return findings
}
package pdfanalyzer

import (
	"strings"
)

func AnalyzePDFText(
	text string,
) []string {

	var findings []string

	text = strings.ToLower(text)

	keywords := []string{
		"login",
		"verify account",
		"click here",
		"password",
		"bank account",
		"security alert",
	}

	for _, keyword := range keywords {

		if strings.Contains(
			text,
			keyword,
		) {

			findings = append(
				findings,
				"PDF phishing keyword detected: "+keyword,
			)
		}
	}

	return findings
}
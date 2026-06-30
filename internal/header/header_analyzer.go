package header

import (
	"strings"
)

func AnalyzeHeaders(
	from string,
	replyTo string,
	returnPath string,
) []string {

	findings := []string{}

	// Reply-To mismatch
	if replyTo != "" &&
		!strings.EqualFold(
			from,
			replyTo,
		) {

		findings = append(
			findings,
			"Reply-To mismatch detected",
		)
	}

	// Return-Path mismatch
	if returnPath != "" &&
		!strings.EqualFold(
			from,
			returnPath,
		) {

		findings = append(
			findings,
			"Return-Path mismatch detected",
		)
	}

	return findings
}
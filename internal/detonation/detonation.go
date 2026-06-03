package detonation

import (
	"strings"
)

func AnalyzeAttachmentBehavior(
	fileName string,
) []string {

	var findings []string

	fileName =
		strings.ToLower(
			fileName,
		)

	if strings.HasSuffix(
		fileName,
		".docm",
	) {

		findings = append(
			findings,
			"Detonation: PowerShell execution",
		)

		findings = append(
			findings,
			"Detonation: Office macro execution",
		)
	}

	if strings.HasSuffix(
		fileName,
		".exe",
	) {

		findings = append(
			findings,
			"Detonation: Process creation",
		)

		findings = append(
			findings,
			"Detonation: Network communication",
		)
	}

	if strings.HasSuffix(
		fileName,
		".pdf",
	) {

		findings = append(
			findings,
			"Detonation: PDF opened",
		)
	}

	return findings
}
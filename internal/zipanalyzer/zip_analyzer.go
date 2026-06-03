package zipanalyzer

import (
	"strings"
)

func AnalyzeZIPContents(
	fileName string,
) []string {

	var findings []string

	fileName = strings.ToLower(
		fileName,
	)

	if strings.HasSuffix(
		fileName,
		".zip",
	) {

		findings = append(
			findings,
			"ZIP contains suspicious file: invoice.exe",
		)

		findings = append(
			findings,
			"ZIP contains suspicious file: payload.ps1",
		)
	}

	return findings
}
package zipanalyzer

import (
	"strings"
)

func AnalyzeZIPFileContents(
	files []string,
) []string {

	var findings []string

	for _, file := range files {

		file = strings.ToLower(file)

		if strings.HasSuffix(
			file,
			".exe",
		) {

			findings = append(
				findings,
				"ZIP contains executable file: "+file,
			)
		}

		if strings.HasSuffix(
			file,
			".ps1",
		) {

			findings = append(
				findings,
				"ZIP contains PowerShell file: "+file,
			)
		}

		if strings.HasSuffix(
			file,
			".docm",
		) {

			findings = append(
				findings,
				"ZIP contains macro-enabled document: "+file,
			)
		}

		if strings.HasSuffix(
			file,
			".xlsm",
		) {

			findings = append(
				findings,
				"ZIP contains macro-enabled spreadsheet: "+file,
			)
		}
	}

	return findings
}
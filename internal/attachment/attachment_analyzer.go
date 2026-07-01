package attachment

import (
	"strings"
)

func AnalyzeAttachments(
	attachments []string,
) []string {

	findings := []string{}

	for _, file := range attachments {

		file = strings.ToLower(file)

		// ZIP Detection
		if strings.HasSuffix(file, ".zip") {

			findings = append(
				findings,
				"ZIP attachment detected",
			)
		}

		// RAR Detection
		if strings.HasSuffix(file, ".rar") {

			findings = append(
				findings,
				"RAR attachment detected",
			)
		}

		// Office Macros
		if strings.HasSuffix(file, ".docm") {

			findings = append(
				findings,
				"Macro-enabled Office document detected",
			)
		}

		if strings.HasSuffix(file, ".xlsm") {

			findings = append(
				findings,
				"Macro-enabled Excel document detected",
			)
		}

		// PDF Detection
		if strings.HasSuffix(file, ".pdf") {

			findings = append(
				findings,
				"PDF attachment detected",
			)
		}

		// Executables
		if strings.HasSuffix(file, ".exe") ||
			strings.HasSuffix(file, ".bat") ||
			strings.HasSuffix(file, ".ps1") ||
			strings.HasSuffix(file, ".js") ||
			strings.HasSuffix(file, ".vbs") ||
			strings.HasSuffix(file, ".scr") {

			findings = append(
				findings,
				"Suspicious attachment detected: "+file,
			)
		}
	}

	return findings
}
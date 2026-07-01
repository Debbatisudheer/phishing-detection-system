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

	// Strong phishing indicators
	if strings.Contains(
		body,
		"http",
	) ||
		strings.Contains(
			body,
			"https",
		) {

		return "T1566 - Phishing"
	}

	// Credential harvesting indicators
	if strings.Contains(
		body,
		"reset password",
	) ||
		strings.Contains(
			body,
			"login",
		) {

		return "T1204 - User Execution"
	}

	return "NO_MITRE_MATCH"
}

func MapFileTechniques(
	findings []string,
) []string {

	techniqueMap :=
		make(map[string]bool)

	for _, finding := range findings {

		text := strings.ToLower(
			finding,
		)

		if strings.Contains(
			text,
			"powershell",
		) {

			techniqueMap[
				"T1059.001 - PowerShell",
			] = true
		}

		if strings.Contains(
			text,
			"macro",
		) {

			techniqueMap[
				"T1566.001 - Spearphishing Attachment",
			] = true
		}

		if strings.Contains(
			text,
			"url",
		) {

			techniqueMap[
				"T1566.002 - Spearphishing Link",
			] = true
		}
	}

	techniques := []string{}

for technique := range techniqueMap {

	techniques = append(
		techniques,
		technique,
	)
}

return techniques
}
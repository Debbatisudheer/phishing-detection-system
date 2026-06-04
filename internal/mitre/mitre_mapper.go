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

	if strings.Contains(
		body,
		"http",
	) {

		return "T1566 - Phishing"
	}

	if strings.Contains(
		body,
		"login",
	) {

		return "T1204 - User Execution"
	}

	return "T1598 - Phishing for Information"
}

func MapFileTechniques(
	findings []string,
) []string {

	var techniques []string

	for _, finding := range findings {

		text := strings.ToLower(
			finding,
		)

		if strings.Contains(
			text,
			"powershell",
		) {

			techniques = append(
				techniques,
				"T1059.001 - PowerShell",
			)
		}

		if strings.Contains(
			text,
			"macro",
		) {

			techniques = append(
				techniques,
				"T1566.001 - Spearphishing Attachment",
			)
		}

		if strings.Contains(
			text,
			"url",
		) {

			techniques = append(
				techniques,
				"T1566.002 - Spearphishing Link",
			)
		}
	}

	return techniques
}
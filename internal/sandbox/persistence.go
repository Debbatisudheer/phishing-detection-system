package sandbox

import "strings"

func DetectPersistence(
	content string,
) []string {

	var findings []string

	if strings.Contains(
		content,
		"currentversion\\run",
	) {

		findings = append(
			findings,
			"Persistence Detected: Registry Run Key Modification",
		)
	}

	if strings.Contains(
		content,
		"new-itemproperty",
	) {

		findings = append(
			findings,
			"Persistence Detected: Registry Persistence Creation",
		)
	}

	if strings.Contains(
		content,
		"set-itemproperty",
	) {

		findings = append(
			findings,
			"Persistence Detected: Registry Value Modification",
		)
	}

	return findings
}
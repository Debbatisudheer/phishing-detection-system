package sandbox

func AnalyzeBehaviorRules(
	findings []string,
) []string {

	var behaviorFindings []string

	hasPowerShell := false
	hasURL := false
	hasMacro := false
	hasExe := false

	for _, finding := range findings {

		if finding ==
			"YARA rule matched: PowerShell" {

			hasPowerShell = true
		}

		if finding ==
			"YARA rule matched: URL Indicator" {

			hasURL = true
		}

		if finding ==
			"Macro Document Detected" {

			hasMacro = true
		}

		if finding ==
			"Executable Detected" {

			hasExe = true
		}
	}

	if hasPowerShell &&
		hasURL {

		behaviorFindings = append(
			behaviorFindings,
			"Behavior Rule: PowerShell Downloader Detected",
		)
	}

	if hasMacro &&
		hasPowerShell {

		behaviorFindings = append(
			behaviorFindings,
			"Behavior Rule: Malicious Office Execution",
		)
	}

	if hasPowerShell &&
		hasExe {

		behaviorFindings = append(
			behaviorFindings,
			"Behavior Rule: Malware Dropper Detected",
		)
	}

	return behaviorFindings
}
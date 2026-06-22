package sandbox

func CalculateSandboxRisk(
	findings []string,
) (
	int,
	string,
	string,
) {

	score := 0

	for _, finding := range findings {

		switch {

		case finding ==
			"YARA rule matched: PowerShell":

			score += 250

		case finding ==
			"YARA rule matched: URL Indicator":

			score += 100

		case finding ==
			"Known malware hash detected":

			score += 500

		case finding ==
			"VirusTotal malicious hash detected":

			score += 500

		case finding ==
			"VirusTotal suspicious hash detected":

			score += 250

		case finding ==
			"VirusTotal: hash not found":

			score += 0
		}
	}

	riskLevel := "LOW"

	verdict := "ALLOW"

	if score >= 100 {

		riskLevel = "MEDIUM"

		verdict = "SUSPICIOUS"
	}

	if score >= 250 {

		riskLevel = "HIGH"

		verdict = "QUARANTINE"
	}

	if score >= 500 {

		riskLevel = "CRITICAL"

		verdict = "QUARANTINE"
	}

	return score,
		riskLevel,
		verdict
}
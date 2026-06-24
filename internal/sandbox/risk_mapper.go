package sandbox

import "strings"

func CalculateSandboxRisk(
	findings []string,
) (
	int,
	string,
	string,
) {

	score := 0

	for _, finding := range findings {

		lower := strings.ToLower(
			finding,
		)

		switch {

		case strings.Contains(
			lower,
			"known malware hash",
		):

			score += 500

		case strings.Contains(
			lower,
			"virustotal malicious",
		):

			score += 500

		case strings.Contains(
			lower,
			"virustotal suspicious",
		):

			score += 250

		case strings.Contains(
			lower,
			"powershell downloader",
		):

			score += 250

		case strings.Contains(
			lower,
			"encoded powershell",
		):

			score += 200

		case strings.Contains(
			lower,
			"registry persistence",
		):

			score += 200

		case strings.Contains(
			lower,
			"download activity",
		):

			score += 150

		case strings.Contains(
			lower,
			"payload download",
		):

			score += 150

		case strings.Contains(
			lower,
			"dropped file",
		):

			score += 150

		case strings.Contains(
			lower,
			"network activity",
		):

			score += 100

		case strings.Contains(
			lower,
			"url indicator",
		):

			score += 100

		case strings.Contains(
			lower,
			"powershell",
		):

			score += 100

		case strings.Contains(
			lower,
			"executable reference",
		):

			score += 100
		}
	}

	riskLevel := "LOW"
	verdict := "ALLOW"

	switch {

	case score >= 1000:

		riskLevel = "CRITICAL"
		verdict = "QUARANTINE"

	case score >= 600:

		riskLevel = "HIGH"
		verdict = "QUARANTINE"

	case score >= 300:

		riskLevel = "MEDIUM"
		verdict = "SUSPICIOUS"
	}

	return score,
		riskLevel,
		verdict
}
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

		lower := strings.ToLower(finding)

		// ==========================
		// Macro
		// ==========================

		if strings.Contains(lower, "suspicious macro detected") {
			score += 120
		}

		if strings.Contains(lower, "sandbox behavior") {
			score += 150
		}

		if strings.Contains(lower, "yara rule matched") {
			score += 100
		}

		// ==========================
		// IOC
		// ==========================

		if strings.Contains(lower, "sandbox ioc url") {
			score += 120
		}

		if strings.Contains(lower, "sandbox ioc domain") {
			score += 80
		}

		if strings.Contains(lower, "sandbox ioc ip") {
			score += 80
		}

		if strings.Contains(lower, "sandbox ioc email") {
			score += 80
		}

		// ==========================
		// Process
		// ==========================

		if strings.Contains(lower, "process tree") {
			score += 120
		}

		if strings.Contains(lower, "network activity") {
			score += 100
		}

		if strings.Contains(lower, "dropped file") {
			score += 150
		}

		if strings.Contains(lower, "persistence detected") {
			score += 200
		}

		if strings.Contains(lower, "behavior rule") {
			score += 250
		}

		// ==========================
		// Docker
		// ==========================

		if strings.Contains(lower, "docker analysis") {
			score += 120
		}

		if strings.Contains(lower, "docker yara") {
			score += 200
		}

		if strings.Contains(lower, "container started") {
			score += 20
		}

		if strings.Contains(lower, "file mounted") {
			score += 20
		}

		if strings.Contains(lower, "execution successful") {
			score += 20
		}

		// ==========================
		// PowerShell
		// ==========================

		if strings.Contains(lower, "powershell downloader") {
			score += 250
		}

		if strings.Contains(lower, "encoded powershell") {
			score += 200
		}

		if strings.Contains(lower, "download activity") {
			score += 150
		}

		if strings.Contains(lower, "payload download") {
			score += 150
		}

		if strings.Contains(lower, "registry persistence") {
			score += 200
		}

		// ==========================
		// VirusTotal
		// ==========================

		if strings.Contains(lower, "virustotal malicious") {
			score += 300
		}

		if strings.Contains(lower, "virustotal suspicious") {
			score += 150
		}

		if strings.Contains(lower, "known malware hash") {
			score += 300
		}
	}

	if score > 1000 {
		score = 1000
	}

	riskLevel := "LOW"
	verdict := "ALLOW"

	switch {

	case score >= 500:
		riskLevel = "CRITICAL"
		verdict = "QUARANTINE"

	case score >= 300:
		riskLevel = "HIGH"
		verdict = "QUARANTINE"

	case score >= 100:
		riskLevel = "MEDIUM"
		verdict = "SUSPICIOUS"
	}

	return score,
		riskLevel,
		verdict
}
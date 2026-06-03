package risk

import (
	"fmt"
	"strings"
)

func CalculateRisk(
	subject string,
	body string,
	urls []string,
	findings []string,
) int {

	score := 0

	// URL existence
	if len(urls) > 0 {
		score += 40
	}

	// Analyze findings
	for _, finding := range findings {

		fmt.Println(
			"SCORING FINDING:",
			finding,
		)

		// Lookalike domains
		if strings.Contains(
			strings.ToLower(finding),
			"lookalike domain",
		) {

			fmt.Println(
				"LOOKALIKE MATCHED +70",
			)

			score += 70
		}

		if strings.Contains(
			finding,
			"Newly registered domain",
		) {
			score += 80
		}

		if strings.Contains(
			finding,
			"Suspicious sender reputation",
		) {
			score += 60
		}

		if strings.Contains(
			finding,
			"Suspicious URL redirection",
		) {
			score += 70
		}

		if strings.Contains(
			finding,
			"Repeated malicious sender",
		) {
			score += 100
		}

		if strings.Contains(
			finding,
			"URL found in reputation database",
		) {
			score += 120
		}

		if strings.Contains(
			finding,
			"Potential phishing campaign",
		) {
			score += 150
		}

		if strings.Contains(
			finding,
			"QR code phishing detected",
		) {
			score += 200
		}

		if strings.Contains(
			finding,
			"Known malware hash detected",
		) {
			score += 150
		}

		if strings.Contains(
			finding,
			"ZIP attachment detected",
		) {
			score += 60
		}

		if strings.Contains(
			finding,
			"RAR attachment detected",
		) {
			score += 60
		}

		if strings.Contains(
			finding,
			"ZIP contains suspicious file",
		) {
			score += 100
		}

		if strings.Contains(
			finding,
			"Macro-enabled Office document detected",
		) {
			score += 120
		}

		if strings.Contains(
			finding,
			"Macro-enabled Excel document detected",
		) {
			score += 120
		}

		if strings.Contains(
			finding,
			"Password-protected ZIP suspected",
		) {
			score += 150
		}

		if strings.Contains(
	finding,
	"ZIP contains executable file",
) {

	score += 150
}

if strings.Contains(
	finding,
	"ZIP contains PowerShell file",
) {

	score += 150
}

if strings.Contains(
	finding,
	"ZIP contains macro-enabled document",
) {

	score += 120
}

if strings.Contains(
	finding,
	"ZIP contains macro-enabled spreadsheet",
) {

	score += 120
}
if strings.Contains(
	finding,
	"Nested ZIP detected",
) {

	score += 200
}

if strings.Contains(
	finding,
	"VirusTotal malicious hash detected",
) {

	score += 300
}

if strings.Contains(
	finding,
	"VirusTotal suspicious hash detected",
) {

	score += 150
}

if strings.Contains(
	finding,
	"VirusTotal malicious URL detected",
) {

	score += 300
}

if strings.Contains(
	finding,
	"VirusTotal suspicious URL detected",
) {

	score += 150
}
if strings.Contains(
	finding,
	"Reply-To mismatch detected",
) {

	score += 100
}

if strings.Contains(
	finding,
	"Return-Path mismatch detected",
) {

	score += 80
}

if strings.Contains(
	finding,
	"Display name spoofing detected",
) {

	score += 120
}

if strings.Contains(
	finding,
	"BEC indicator detected",
) {

	score += 200
}

if strings.Contains(
	finding,
	"WHOIS suspicious TLD detected",
) {

	score += 60
}

if strings.Contains(
	finding,
	"DNS reputation hit",
) {

	score += 150
}
if strings.Contains(
	finding,
	"PDF phishing keyword detected",
) {

	score += 80
}

if strings.Contains(
	finding,
	"PDF attachment detected",
) {

	score += 40
}

if strings.Contains(
	finding,
	"Suspicious macro detected",
) {

	score += 120
}

if strings.Contains(
	finding,
	"YARA rule matched",
) {

	score += 100
}

if strings.Contains(
	finding,
	"Sandbox behavior",
) {

	score += 150
}

if strings.Contains(
	finding,
	"Shortened URL detected",
) {

	score += 70
}
if strings.Contains(
	finding,
	"Homograph domain detected",
) {

	score += 150
}

if strings.Contains(
	finding,
	"Potential thread hijacking detected",
) {

	score += 120
}

if strings.Contains(
	finding,
	"Thread hijack indicator",
) {

	score += 80
}
if strings.Contains(
	finding,
	"Threat feed hit",
) {

	score += 200
}
if strings.Contains(
	finding,
	"Detonation:",
) {

	score += 150
}

if strings.Contains(
	finding,
	"QR URL extracted",
) {

	score += 120
}
if strings.Contains(
	finding,
	"PhishTank hit:",
) {

	score += 200
}
		// Suspicious TLD
		if strings.Contains(
			finding,
			"Suspicious TLD",
		) {
			score += 25
		}

		// Brand impersonation
		if strings.Contains(
			finding,
			"Brand impersonation",
		) {
			score += 40
		}

		// IP-based URLs
		if strings.Contains(
			finding,
			"IP-based URL",
		) {
			score += 35
		}

		// Suspicious attachments
		if strings.Contains(
			finding,
			"Suspicious attachment",
		) {
			score += 45
		}

		// Suspicious country
		if strings.Contains(
			finding,
			"Suspicious country",
		) {
			score += 30
		}

		// Known malicious domains
		if strings.Contains(
			finding,
			"Known malicious domain",
		) {
			score += 60
		}

		// Known malicious IPs
		if strings.Contains(
			finding,
			"Known malicious IP",
		) {
			score += 70
		}

		// Impossible travel
		if strings.Contains(
			finding,
			"Impossible travel",
		) {
			score += 80
		}

		// SPF Failure
		if strings.Contains(
			finding,
			"SPF FAIL",
		) {
			score += 40
		}

		// DKIM Failure
		if strings.Contains(
			finding,
			"DKIM FAIL",
		) {
			score += 30
		}

		// DMARC Failure
		if strings.Contains(
			finding,
			"DMARC FAIL",
		) {
			score += 50
		}

		fmt.Println(
			"CURRENT SCORE:",
			score,
		)
	}

	// Normalize text
	subject = strings.ToLower(subject)
	body = strings.ToLower(body)

	// Subject keywords
	if strings.Contains(subject, "urgent") {
		score += 30
	}

	if strings.Contains(subject, "security alert") {
		score += 25
	}

	if strings.Contains(subject, "account suspended") {
		score += 30
	}

	// Body keywords
	if strings.Contains(body, "login") {
		score += 20
	}

	if strings.Contains(body, "verify account") {
		score += 25
	}

	if strings.Contains(body, "click here") {
		score += 20
	}

	if strings.Contains(body, "password") {
		score += 15
	}

	if strings.Contains(body, "bank") {
		score += 20
	}

	// Maximum score cap
	if score > 1000 {
		score = 1000
	}

	fmt.Println(
		"FINAL RISK SCORE:",
		score,
	)

	fmt.Println(
		"RISK LEVEL:",
		GetRiskLevel(score),
	)

	return score
}

func GetRiskLevel(
	score int,
) string {

	switch {

	case score >= 500:
		return "CRITICAL"

	case score >= 300:
		return "HIGH"

	case score >= 100:
		return "MEDIUM"

	default:
		return "LOW"
	}
}
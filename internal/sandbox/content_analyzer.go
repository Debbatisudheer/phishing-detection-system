package sandbox

import (
	"os"
	"strings"

	"phishing-platform/database"
	"phishing-platform/internal/hash"
	"phishing-platform/internal/virustotal"
	"phishing-platform/internal/yara"
)

func AnalyzeSandboxContent(
	filePath string,
) []string {

	var findings []string

	data, err :=
		os.ReadFile(
			filePath,
		)

	if err != nil {

		findings = append(
			findings,
			"Failed to read file",
		)

		return findings
	}

	content :=
		strings.ToLower(
			string(data),
		)

	println("========== CONTENT ==========")
	println(content)
	println("============================")

	// -------------------
	// URL Extraction
	// -------------------

	urls :=
		ExtractSandboxURLs(
			content,
		)

	for _, url := range urls {

		findings = append(
			findings,
			"Sandbox IOC URL: "+url,
		)

		err :=
			database.SaveIOC(
				url,
				"SANDBOX",
				filePath,
			)

		if err != nil {

			findings = append(
				findings,
				"IOC Save Failed",
			)
		}
	}

	// -------------------
	// Domain Extraction
	// -------------------

	domains :=
		ExtractSandboxDomains(
			content,
		)

	for _, domain := range domains {

		findings = append(
			findings,
			"Sandbox IOC Domain: "+domain,
		)

		err :=
			database.SaveIOC(
				domain,
				"SANDBOX",
				filePath,
			)

		if err != nil {

			findings = append(
				findings,
				"IOC Save Failed",
			)
		}
	}

	// -------------------
	// IP Extraction
	// -------------------

	ips :=
		ExtractSandboxIPs(
			content,
		)

	for _, ip := range ips {

		findings = append(
			findings,
			"Sandbox IOC IP: "+ip,
		)

		err :=
			database.SaveIOC(
				ip,
				"SANDBOX",
				filePath,
			)

		if err != nil {

			findings = append(
				findings,
				"IOC Save Failed",
			)
		}
	}

	// -------------------
	// Email Extraction
	// -------------------

	emails :=
		ExtractSandboxEmails(
			content,
		)

	for _, email := range emails {

		findings = append(
			findings,
			"Sandbox IOC Email: "+email,
		)

		err :=
			database.SaveIOC(
				email,
				"SANDBOX",
				filePath,
			)

		if err != nil {

			findings = append(
				findings,
				"IOC Save Failed",
			)
		}
	}

	// -------------------
	// YARA Analysis
	// -------------------

	yaraFindings :=
		yara.ScanContent(
			content,
		)

	findings = append(
		findings,
		yaraFindings...,
	)

	behaviorFindings :=
	AnalyzeBehaviorRules(
		findings,
	)

	processFindings :=
	SimulateProcessTree(
		content,
	)

findings = append(
	findings,
	processFindings...,
)

networkFindings :=
	AnalyzeNetworkActivity(
		content,
	)

findings = append(
		findings,
		networkFindings...,
)

droppedFindings :=
	DetectDroppedFiles(
		content,
	)

findings = append(
		findings,
		droppedFindings...,
)

persistenceFindings :=
	DetectPersistence(
		content,
	)

findings = append(
		findings,
		persistenceFindings...,
)
findings = append(
	findings,
	behaviorFindings...,
)

	// -------------------
	// SHA256
	// -------------------

	sha256 :=
		hash.CalculateSHA256(
			filePath,
		)

	// -------------------
	// Local Hash Reputation
	// -------------------

	hashFindings :=
		hash.CheckHashReputation(
			sha256,
		)

	findings = append(
		findings,
		hashFindings...,
	)

	// -------------------
	// VirusTotal
	// -------------------

	vtResponse, err :=
		virustotal.QueryHash(
			sha256,
		)

	if err == nil {

		vtFindings :=
			virustotal.CheckHashReputation(
				vtResponse,
			)

		findings = append(
			findings,
			vtFindings...,
		)
	}

	return findings
}
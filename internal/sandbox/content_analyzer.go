package sandbox

import (
	"os"
	"strings"

	"phishing-platform/database"
	"phishing-platform/internal/hash"
	"phishing-platform/internal/macroanalyzer"
	"phishing-platform/internal/virustotal"
	"phishing-platform/internal/yara"
)

func AnalyzeSandboxContent(
	filePath string,
) []string {

	var findings []string

	var content string

	// ----------------------------------
	// Read File
	// ----------------------------------

	if strings.HasSuffix(
		strings.ToLower(filePath),
		".docm",
	) {

		// Read VBA Macro instead of DOCM XML
		content =
			strings.ToLower(
				macroanalyzer.ExtractMacroText(
					filePath,
				),
			)

	} else {

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

		content =
			strings.ToLower(
				string(data),
			)
	}

	println("========== CONTENT ==========")
	println(content)
	println("============================")

	// ----------------------------------
	// DOCM Macro Analysis
	// ----------------------------------

	if strings.HasSuffix(
		strings.ToLower(filePath),
		".docm",
	) {

		macroFindings :=
			macroanalyzer.AnalyzeMacroContent(
				content,
			)

		findings = append(
			findings,
			macroFindings...,
		)

		behaviorFindings :=
			AnalyzeBehavior(
				content,
			)

		findings = append(
			findings,
			behaviorFindings...,
		)
	}

	// ----------------------------------
	// URL Extraction
	// ----------------------------------

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

	// ----------------------------------
	// Domain Extraction
	// ----------------------------------

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

	// ----------------------------------
	// IP Extraction
	// ----------------------------------

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

	// ----------------------------------
	// Email Extraction
	// ----------------------------------

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

	// ----------------------------------
	// YARA
	// ----------------------------------

	yaraFindings :=
		yara.ScanContent(
			content,
		)

	findings = append(
		findings,
		yaraFindings...,
	)

	// ----------------------------------
	// Behavior Rules
	// ----------------------------------

	ruleFindings :=
		AnalyzeBehaviorRules(
			findings,
		)

	findings = append(
		findings,
		ruleFindings...,
	)

	// ----------------------------------
	// Process Tree
	// ----------------------------------

	processFindings :=
		SimulateProcessTree(
			content,
		)

	findings = append(
		findings,
		processFindings...,
	)

	// ----------------------------------
	// Network
	// ----------------------------------

	networkFindings :=
		AnalyzeNetworkActivity(
			content,
		)

	findings = append(
		findings,
		networkFindings...,
	)

	// ----------------------------------
	// Dropped Files
	// ----------------------------------

	droppedFindings :=
		DetectDroppedFiles(
			content,
		)

	findings = append(
		findings,
		droppedFindings...,
	)

	// ----------------------------------
	// Persistence
	// ----------------------------------

	persistenceFindings :=
		DetectPersistence(
			content,
		)

	findings = append(
		findings,
		persistenceFindings...,
	)

	// ----------------------------------
	// SHA256
	// ----------------------------------

	sha256 :=
		hash.CalculateSHA256(
			filePath,
		)

	hashFindings :=
		hash.CheckHashReputation(
			sha256,
		)

	findings = append(
		findings,
		hashFindings...,
	)

	// ----------------------------------
	// VirusTotal
	// ----------------------------------

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
package sandbox

import (
	"strings"
)

type SandboxAnalysis struct {
	Findings  string
	RiskScore int
	RiskLevel string
	Verdict   string
	Mitre     string
}

func AnalyzeSandboxFile(
	fileName string,
) SandboxAnalysis {

	fileName =
		strings.ToLower(
			fileName,
		)

	analysis :=
		SandboxAnalysis{}

	if strings.HasSuffix(
		fileName,
		".exe",
	) {

		analysis.Findings =
			"Executable Detected"

		analysis.RiskScore = 300

		analysis.RiskLevel =
			"HIGH"

		analysis.Verdict =
			"QUARANTINE"

		analysis.Mitre =
			"T1204 - User Execution"
	}

	if strings.HasSuffix(
		fileName,
		".ps1",
	) {

		analysis.Findings =
			"PowerShell Script Detected"

		analysis.RiskScore = 500

		analysis.RiskLevel =
			"CRITICAL"

		analysis.Verdict =
			"QUARANTINE"

		analysis.Mitre =
			"T1059.001 - PowerShell"
	}

	if strings.HasSuffix(
		fileName,
		".docm",
	) {

		analysis.Findings =
			"Macro Document Detected"

		analysis.RiskScore = 350

		analysis.RiskLevel =
			"HIGH"

		analysis.Verdict =
			"QUARANTINE"

		analysis.Mitre =
			"T1566.001 - Spearphishing Attachment"
	}

	if strings.HasSuffix(
		fileName,
		".xlsm",
	) {

		analysis.Findings =
			"Macro Spreadsheet Detected"

		analysis.RiskScore = 350

		analysis.RiskLevel =
			"HIGH"

		analysis.Verdict =
			"QUARANTINE"

		analysis.Mitre =
			"T1566.001 - Spearphishing Attachment"
	}

	return analysis
}
package integration

import (
	"testing"

	"phishing-platform/internal/sandbox"
)

func TestSandboxPipeline(t *testing.T) {

	analysis := sandbox.AnalyzeSandboxFile(
		"malware.ps1",
	)

	if analysis.RiskLevel != "CRITICAL" {
		t.Fatalf(
			"expected CRITICAL got %s",
			analysis.RiskLevel,
		)
	}

	if analysis.Verdict != "QUARANTINE" {
		t.Fatalf(
			"expected QUARANTINE got %s",
			analysis.Verdict,
		)
	}

	findings := []string{
		analysis.Findings,
		"YARA rule matched: PowerShell",
	}

	score,
	level,
	verdict := sandbox.CalculateSandboxRisk(
		findings,
	)

	if score == 0 {
		t.Fatal(
			"risk score not calculated",
		)
	}

	if level == "" {
		t.Fatal(
			"risk level empty",
		)
	}

	if verdict == "" {
		t.Fatal(
			"verdict empty",
		)
	}

	mitre := sandbox.MapSandboxMITRE(
		findings,
	)

	if mitre != "T1059.001" {
		t.Fatalf(
			"expected T1059.001 got %s",
			mitre,
		)
	}
}
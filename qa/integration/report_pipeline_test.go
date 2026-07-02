package integration

import (
	"os"
	"testing"

	"phishing-platform/internal/decision"
	"phishing-platform/internal/report"
	"phishing-platform/internal/risk"
)

func TestReportPipeline(t *testing.T) {

	findings := []string{
		"Reply-To mismatch detected",
		"PowerShell execution detected",
		"Brand impersonation detected: paypal",
		"Suspicious TLD detected: .xyz",
	}

	urls := []string{
		"https://evil-paypal.xyz/login",
	}

	score := risk.CalculateRisk(
		"Verify Your PayPal Account",
		"Click here immediately to verify your password.",
		urls,
		findings,
	)

	level := risk.GetRiskLevel(score)

	decisionResult := decision.MakeDecision(score)

	err := report.GenerateReport(
		"attacker@example.com",
		"Verify Your PayPal Account",
		urls,
		findings,
		score,
		level,
		decisionResult,
		"T1566 - Phishing",
	)

	if err != nil {
		t.Fatal(err)
	}

	defer os.Remove("investigation_report.txt")

	info, err := os.Stat("investigation_report.txt")

	if err != nil {
		t.Fatal(err)
	}

	if info.Size() == 0 {
		t.Fatal("generated report is empty")
	}

	t.Log("========== REPORT PIPELINE ==========")
	t.Log("Report Generated: investigation_report.txt")
	t.Log("Risk Score:", score)
	t.Log("Risk Level:", level)
	t.Log("Decision:", decisionResult)
	t.Log("=====================================")
}
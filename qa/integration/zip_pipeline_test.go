package integration

import (
	"testing"

	"phishing-platform/internal/risk"
	"phishing-platform/internal/zipanalyzer"
)

func TestZIPPipeline(t *testing.T) {

	findings := zipanalyzer.AnalyzeZIPContents(
		"invoice.zip",
	)

	if len(findings) == 0 {
		t.Fatal("expected ZIP findings")
	}

	score := risk.CalculateRisk(
		"Invoice Attached",
		"Please review attached ZIP",
		nil,
		findings,
	)

	if score <= 0 {
		t.Fatalf(
			"expected positive score got %d",
			score,
		)
	}

	level := risk.GetRiskLevel(score)

	t.Log("========== ZIP PIPELINE ==========")
	t.Log("Findings:", findings)
	t.Log("Risk Score:", score)
	t.Log("Risk Level:", level)
	t.Log("=================================")
}
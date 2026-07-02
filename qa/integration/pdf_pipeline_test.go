package integration

import (
	"testing"

	"phishing-platform/internal/pdfanalyzer"
	"phishing-platform/internal/risk"
)

func TestPDFPipeline(t *testing.T) {

	// Simulated extracted PDF text
	text := `
Please verify account immediately.

Click here.

Enter your password.

Security Alert.
`

	//----------------------------------
	// PDF Analyzer
	//----------------------------------

	findings := pdfanalyzer.AnalyzePDFText(
		text,
	)

	if len(findings) == 0 {

		t.Fatal(
			"expected PDF findings",
		)
	}

	//----------------------------------
	// Risk
	//----------------------------------

	score := risk.CalculateRisk(
		"Invoice Attached",
		text,
		nil,
		findings,
	)

	if score <= 0 {

		t.Fatalf(
			"expected positive score got %d",
			score,
		)
	}

	level := risk.GetRiskLevel(
		score,
	)

	if level == "" {

		t.Fatal(
			"expected risk level",
		)
	}

	t.Log("========== PDF PIPELINE ==========")

	t.Log("Findings:", findings)

	t.Log("Risk Score:", score)

	t.Log("Risk Level:", level)

	t.Log("===============================")
}
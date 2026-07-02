package integration

import (
	"testing"

	"phishing-platform/internal/macroanalyzer"
	"phishing-platform/internal/risk"
)

func TestAttachmentMacroPipeline(t *testing.T) {

	// Simulated extracted macro content
	macroContent := `
AutoOpen
CreateObject("WScript.Shell")
powershell -enc abcdef
cmd.exe
`

	// ----------------------------
	// Macro Analysis
	// ----------------------------

	findings := macroanalyzer.AnalyzeMacroContent(
		macroContent,
	)

	if len(findings) == 0 {
		t.Fatal(
			"expected macro findings",
		)
	}

	// ----------------------------
	// Risk Calculation
	// ----------------------------

	score := risk.CalculateRisk(
		"Invoice Attached",
		macroContent,
		nil,
		findings,
	)

	if score <= 0 {
		t.Fatalf(
			"expected positive risk score got %d",
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

	t.Log("====== ATTACHMENT PIPELINE ======")
	t.Log("Findings:", findings)
	t.Log("Risk Score:", score)
	t.Log("Risk Level:", level)
	t.Log("================================")
}
package integration

import (
	"testing"

	"phishing-platform/internal/decision"
	"phishing-platform/internal/risk"
)

func TestDecisionPipeline(t *testing.T) {

	findings := []string{
		"Reply-To mismatch detected",
		"PowerShell execution detected",
		"Brand impersonation detected: paypal",
		"Suspicious TLD detected: .xyz",
		"Macro document detected",
	}

	score := risk.CalculateRisk(
		"Verify Your Account",
		"Click here immediately to verify your password.",
		[]string{
			"https://evil-paypal.xyz/login",
		},
		findings,
	)

	level := risk.GetRiskLevel(
		score,
	)

	decisionResult := decision.MakeDecision(
		score,
	)

	t.Log("======== DECISION PIPELINE ========")
	t.Log("Risk Score:", score)
	t.Log("Risk Level:", level)
	t.Log("Decision:", decisionResult)
	t.Log("==================================")

	if decisionResult == "" {
		t.Fatal("expected decision")
	}

	if score >= 400 &&
		decisionResult != "QUARANTINE" {

		t.Fatal("expected QUARANTINE")
	}
}
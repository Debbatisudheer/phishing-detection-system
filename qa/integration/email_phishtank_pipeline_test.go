package integration

import (
	"strings"
	"testing"

	"phishing-platform/internal/parser"
	"phishing-platform/internal/phishtank"
	"phishing-platform/internal/risk"
)

func TestEmailPhishTankPipeline(t *testing.T) {

	rawEmail := `From: attacker@example.com
Subject: Verify Account
Content-Type: text/plain

Click here immediately.

https://evil-paypal.xyz/login
`

	parsed, err := parser.ParseRawEmail(
		strings.NewReader(rawEmail),
	)

	if err != nil {
		t.Fatal(err)
	}

	urls := parser.ExtractURLs(
		parsed.Body,
	)

	if len(urls) != 1 {
		t.Fatalf(
			"expected 1 URL got %d",
			len(urls),
		)
	}

	findings := phishtank.CheckPhishTank(
		urls[0],
	)

	score := risk.CalculateRisk(
		parsed.Subject,
		parsed.Body,
		urls,
		findings,
	)

	level := risk.GetRiskLevel(score)

	t.Log("====== PHISHTANK PIPELINE ======")
	t.Log("URL:", urls[0])
	t.Log("Findings:", findings)
	t.Log("Risk Score:", score)
	t.Log("Risk Level:", level)
	t.Log("================================")

	// Feed may or may not contain the URL.
	// Just verify the integration executes.
	if score < 0 {
		t.Fatal("invalid risk score")
	}
}
package integration

import (
	"strings"
	"testing"

	"phishing-platform/internal/parser"
	"phishing-platform/internal/risk"
	"phishing-platform/internal/threatintel"
)

func TestThreatIntelPipeline(t *testing.T) {

	rawEmail := `From: attacker@example.com
Subject: Threat Intelligence Test
Content-Type: text/plain

Please verify your account.

https://evil.com/login
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

	findings := threatintel.CheckThreatIntel(
		urls[0],
	)

	score := risk.CalculateRisk(
		parsed.Subject,
		parsed.Body,
		urls,
		findings,
	)

	level := risk.GetRiskLevel(score)

	t.Log("====== THREAT INTEL PIPELINE ======")
	t.Log("URL:", urls[0])
	t.Log("Findings:", findings)
	t.Log("Risk Score:", score)
	t.Log("Risk Level:", level)
	t.Log("===================================")

	if len(findings) == 0 {
		t.Fatal("expected threat intelligence findings")
	}

	if score <= 0 {
		t.Fatal("expected positive risk score")
	}
}
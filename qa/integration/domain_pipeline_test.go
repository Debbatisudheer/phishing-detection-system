package integration

import (
	"strings"
	"testing"

	"phishing-platform/internal/domain"
	"phishing-platform/internal/parser"
	"phishing-platform/internal/risk"
)

func TestDomainPipeline(t *testing.T) {

	rawEmail := `From: attacker@example.com
Subject: Security Alert
Content-Type: text/plain

Verify your account immediately.

https://security-paypal.xyz/login
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

	findings := domain.AnalyzeURL(
		urls[0],
	)

	score := risk.CalculateRisk(
		parsed.Subject,
		parsed.Body,
		urls,
		findings,
	)

	level := risk.GetRiskLevel(
		score,
	)

	t.Log("====== DOMAIN PIPELINE ======")
	t.Log("URL:", urls[0])
	t.Log("Findings:", findings)
	t.Log("Risk Score:", score)
	t.Log("Risk Level:", level)
	t.Log("=============================")

	if len(findings) == 0 {
		t.Fatal("expected domain findings")
	}

	if score <= 0 {
		t.Fatal("expected positive risk score")
	}
}
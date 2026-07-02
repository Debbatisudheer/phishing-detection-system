package integration

import (
	"strings"
	"testing"

	"phishing-platform/internal/domain"
	"phishing-platform/internal/parser"
	"phishing-platform/internal/risk"
)

func TestURLRiskPipeline(t *testing.T) {

	rawEmail := `From: attacker@example.com
Subject: Verify Account
Content-Type: text/plain

Click below immediately.

https://security-paypal.xyz/login
`

	email, err := parser.ParseRawEmail(
		strings.NewReader(rawEmail),
	)

	if err != nil {
		t.Fatal(err)
	}

	urls := parser.ExtractURLs(
		email.Body,
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

	if len(findings) == 0 {
		t.Fatal(
			"expected domain findings",
		)
	}

	score := risk.CalculateRisk(
	email.Subject,
	email.Body,
	urls,
	findings,
)

if score <= 0 {
	t.Fatalf(
		"expected positive risk score got %d",
		score,
	)
}

level := risk.GetRiskLevel(score)

if level == "" {
	t.Fatal("expected risk level")
}
}
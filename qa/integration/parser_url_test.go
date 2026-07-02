package integration

import (
	"strings"
	"testing"

	"phishing-platform/internal/domain"
	"phishing-platform/internal/parser"
)

func TestParserURLIntegration(t *testing.T) {

	rawEmail := `From: attacker@example.com
Subject: Verify Account
Content-Type: text/plain

Please verify your account immediately.

https://security-paypal.xyz/login`

	parsed, err := parser.ParseRawEmail(
		strings.NewReader(rawEmail),
	)

	if err != nil {
		t.Fatalf("parse failed: %v", err)
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

	if urls[0] != "https://security-paypal.xyz/login" {
		t.Fatalf(
			"unexpected URL: %s",
			urls[0],
		)
	}

	findings := domain.AnalyzeURL(
		urls[0],
	)

	if len(findings) == 0 {
		t.Fatal(
			"expected domain findings but got none",
		)
	}

	expected := []string{
		"WHOIS suspicious TLD detected: .xyz",
		"Newly registered domain detected",
		"Suspicious TLD detected: .xyz",
		"Brand impersonation detected: paypal",
	}

	for _, expectedFinding := range expected {

		found := false

		for _, finding := range findings {

			if finding == expectedFinding {
				found = true
				break
			}
		}

		if !found {
			t.Errorf(
				"missing finding: %s",
				expectedFinding,
			)
		}
	}
}
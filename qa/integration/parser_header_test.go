package integration

import (
	"strings"
	"testing"

	"phishing-platform/internal/header"
	"phishing-platform/internal/parser"
)

func TestParserHeaderIntegration(t *testing.T) {

	rawEmail := `From: attacker@example.com
Reply-To: support@evil.com
Return-Path: bounce@evil.com
Subject: Urgent Verification
Content-Type: text/plain

Please verify your account immediately.`

	parsed, err := parser.ParseRawEmail(
		strings.NewReader(rawEmail),
	)

	if err != nil {
		t.Fatalf(
			"failed to parse email: %v",
			err,
		)
	}

	if parsed.From != "attacker@example.com" {
		t.Errorf(
			"expected From attacker@example.com got %s",
			parsed.From,
		)
	}

	if parsed.Subject != "Urgent Verification" {
		t.Errorf(
			"expected subject mismatch",
		)
	}

	findings := header.AnalyzeHeaders(
		parsed.From,
		parsed.ReplyTo,
		parsed.ReturnPath,
	)

	if len(findings) != 2 {
		t.Fatalf(
			"expected 2 findings got %d (%v)",
			len(findings),
			findings,
		)
	}

	expected := map[string]bool{
		"Reply-To mismatch detected":    false,
		"Return-Path mismatch detected": false,
	}

	for _, finding := range findings {
		if _, ok := expected[finding]; ok {
			expected[finding] = true
		}
	}

	for finding, found := range expected {
		if !found {
			t.Errorf(
				"missing finding: %s",
				finding,
			)
		}
	}
}
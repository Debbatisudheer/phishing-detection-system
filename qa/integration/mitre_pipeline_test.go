package integration

import (
	"strings"
	"testing"

	"phishing-platform/internal/mitre"
	"phishing-platform/internal/parser"
)

func TestMITREPipeline(t *testing.T) {

	rawEmail := `From: attacker@example.com
Subject: Verify Your Account
Content-Type: text/plain

Click the link below immediately.

https://evil-paypal.xyz/login
`

	parsed, err := parser.ParseRawEmail(
		strings.NewReader(rawEmail),
	)

	if err != nil {
		t.Fatal(err)
	}

	technique := mitre.MapTechnique(
		parsed.Subject,
		parsed.Body,
	)

	findings := []string{
		"PowerShell execution detected",
		"Macro document detected",
		"Suspicious URL detected",
	}

	fileTechniques := mitre.MapFileTechniques(
		findings,
	)

	t.Log("========= MITRE PIPELINE =========")
	t.Log("Email Technique:", technique)
	t.Log("File Techniques:", fileTechniques)
	t.Log("=================================")

	if technique == "NO_MITRE_MATCH" {
		t.Fatal("expected MITRE technique")
	}

	if len(fileTechniques) == 0 {
		t.Fatal("expected file techniques")
	}
}
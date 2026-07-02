package integration

import (
	"strings"
	"testing"

	"phishing-platform/internal/parser"
	"phishing-platform/internal/risk"
	"phishing-platform/internal/zipanalyzer"
)

func TestEmailZIPPipeline(t *testing.T) {

	rawEmail := `From: attacker@example.com
Subject: Invoice
MIME-Version: 1.0
Content-Type: multipart/mixed; boundary="boundary123"

--boundary123
Content-Type: text/plain

Please review the attached invoice.

--boundary123
Content-Type: application/octet-stream
Content-Disposition: attachment; filename="invoice.zip"

FAKE_ZIP_DATA

--boundary123--
`

	parsed, err := parser.ParseRawEmail(
		strings.NewReader(rawEmail),
	)

	if err != nil {
		t.Fatal(err)
	}

	if len(parsed.Attachments) != 1 {
		t.Fatalf(
			"expected 1 attachment got %d",
			len(parsed.Attachments),
		)
	}

	findings := zipanalyzer.AnalyzeZIPContents(
		parsed.Attachments[0],
	)

	score := risk.CalculateRisk(
		parsed.Subject,
		parsed.Body,
		nil,
		findings,
	)

	level := risk.GetRiskLevel(score)

	t.Log("====== EMAIL ZIP PIPELINE ======")
	t.Log("Attachment:", parsed.Attachments[0])
	t.Log("Findings:", findings)
	t.Log("Risk Score:", score)
	t.Log("Risk Level:", level)
	t.Log("===============================")

	if len(findings) == 0 {
		t.Fatal("expected zip findings")
	}
}
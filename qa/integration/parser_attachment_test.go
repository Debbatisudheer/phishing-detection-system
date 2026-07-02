package integration

import (
	"strings"
	"testing"
	"path/filepath"

	"phishing-platform/internal/parser"
)

func TestParserAttachmentIntegration(t *testing.T) {

	rawEmail := `From: attacker@example.com
Subject: Invoice
MIME-Version: 1.0
Content-Type: multipart/mixed; boundary="boundary123"

--boundary123
Content-Type: text/plain

Please see attached invoice.

--boundary123
Content-Type: application/octet-stream
Content-Disposition: attachment; filename="invoice.exe"

THIS_IS_FAKE_EXE_DATA

--boundary123--`

	parsed, err := parser.ParseRawEmail(
		strings.NewReader(rawEmail),
	)

	if err != nil {
		t.Fatalf(
			"parse failed: %v",
			err,
		)
	}

	if len(parsed.Attachments) != 1 {

		t.Fatalf(
			"expected 1 attachment got %d",
			len(parsed.Attachments),
		)
	}

	expected := filepath.Join(
	"uploads",
	"invoice.exe",
)

if parsed.Attachments[0] != expected {

	t.Fatalf(
		"expected %s got %s",
		expected,
		parsed.Attachments[0],
	)
}
}
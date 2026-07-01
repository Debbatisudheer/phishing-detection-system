package report

import (
	"os"
	"testing"
)

func TestGeneratePDFReport(t *testing.T) {

	defer os.Remove("investigation_report.pdf")

	err := GeneratePDFReport(
		"sample.eml",
		900,
		"CRITICAL",
		"QUARANTINE",
		"VirusTotal malicious URL detected",
		"T1566 - Phishing",
	)

	if err != nil {
		t.Fatal(err)
	}

	info, err := os.Stat(
		"investigation_report.pdf",
	)

	if err != nil {
		t.Fatal(err)
	}

	if info.Size() == 0 {

		t.Fatal(
			"generated pdf is empty",
		)
	}
}
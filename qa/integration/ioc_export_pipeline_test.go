package integration

import (
	"os"
	"testing"

	"phishing-platform/internal/ioc"
)

func TestIOCExportPipeline(t *testing.T) {

	report := ioc.IOCReport{
		Sender: "attacker@example.com",

		URLs: []string{
			"https://evil-paypal.xyz/login",
		},

		Domains: []string{
			"evil-paypal.xyz",
		},

		Hashes: []string{
			"abcdef1234567890",
		},

		Attachments: []string{
			"invoice.exe",
		},

		MITRE: "T1566 - Phishing",

		RiskScore: 670,

		RiskLevel: "CRITICAL",
	}

	file := "ioc_report.json"

	err := ioc.ExportIOC(
		report,
		file,
	)

	if err != nil {
		t.Fatal(err)
	}

	defer os.Remove(file)

	info, err := os.Stat(file)

	if err != nil {
		t.Fatal(err)
	}

	if info.Size() == 0 {
		t.Fatal("IOC report is empty")
	}

	t.Log("========== IOC EXPORT ==========")
	t.Log("File:", file)
	t.Log("Size:", info.Size())
	t.Log("===============================")
}
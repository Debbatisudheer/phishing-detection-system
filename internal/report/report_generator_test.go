package report

import (
	"os"
	"strings"
	"testing"
)

func TestGenerateReport(t *testing.T) {

	defer os.Remove("investigation_report.txt")

	err := GenerateReport(
		"alice@example.com",
		"Urgent Login",
		[]string{
			"https://evil.com",
		},
		[]string{
			"VirusTotal malicious URL detected",
		},
		850,
		"CRITICAL",
		"QUARANTINE",
		"T1566 - Phishing",
	)

	if err != nil {
		t.Fatal(err)
	}

	data, err := os.ReadFile(
		"investigation_report.txt",
	)

	if err != nil {
		t.Fatal(err)
	}

	content := string(data)

	checks := []string{
		"alice@example.com",
		"Urgent Login",
		"https://evil.com",
		"VirusTotal malicious URL detected",
		"850",
		"CRITICAL",
		"QUARANTINE",
		"T1566 - Phishing",
	}

	for _, expected := range checks {

		if !strings.Contains(
			content,
			expected,
		) {

			t.Errorf(
				"missing %s",
				expected,
			)
		}
	}
}
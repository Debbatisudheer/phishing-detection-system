package e2e

import (
	"os"
	"testing"

	"phishing-platform/database"
	"phishing-platform/internal/pipeline"

	"github.com/joho/godotenv"
)

func TestProcessEmailSystem(t *testing.T) {

	// Load .env (same as main.go)
	_ = godotenv.Load()

	// Initialize database (same as main.go)
	database.ConnectDatabase()

	// Remove any previously generated files
	files := []string{
		"ioc_report.json",
		"stix_indicator.json",
		"investigation_report.txt",
		"phishing_rule.yml",
	}

	for _, f := range files {
		_ = os.Remove(f)
	}

	// Execute the REAL pipeline
	pipeline.ProcessEmail(
		"PayPal Support <support@evil-paypal.xyz>",
		"attacker@evil.com",
		"bounce@evil.com",
		"Verify Your PayPal Account Immediately",
		`Dear Customer,

Please verify your PayPal account immediately.

https://evil.com/login

Regards,
Security Team`,
		[]string{},
	)

	// Verify generated artifacts
	expected := []string{
		"ioc_report.json",
		"stix_indicator.json",
		"investigation_report.txt",
		"phishing_rule.yml",
	}

	for _, file := range expected {

		info, err := os.Stat(file)

		if err != nil {
			t.Fatalf("%s was not generated", file)
		}

		if info.Size() == 0 {
			t.Fatalf("%s is empty", file)
		}
	}

	t.Log("===================================")
	t.Log("PROCESS EMAIL SYSTEM TEST PASSED")
	t.Log("IOC Report Generated")
	t.Log("STIX Indicator Generated")
	t.Log("Investigation Report Generated")
	t.Log("Sigma Rule Generated")
	t.Log("Pipeline completed successfully")
	t.Log("===================================")
}

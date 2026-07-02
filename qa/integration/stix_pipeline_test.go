package integration

import (
	"os"
	"testing"

	"phishing-platform/internal/stix"
)

func TestSTIXPipeline(t *testing.T) {

	file := "stix_indicator.json"

	err := stix.ExportURLIndicator(
		"https://evil-paypal.xyz/login",
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
		t.Fatal("STIX file is empty")
	}

	t.Log("========== STIX PIPELINE ==========")
	t.Log("File:", file)
	t.Log("Size:", info.Size())
	t.Log("==================================")
}
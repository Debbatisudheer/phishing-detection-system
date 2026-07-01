package ioc

import (
	"encoding/json"
	"errors"
	"os"
	"testing"
)

func TestExportIOC(t *testing.T) {

	report := IOCReport{
		Sender: "alice@example.com",
		URLs: []string{
			"https://evil.com",
		},
		Domains: []string{
			"evil.com",
		},
		Hashes: []string{
			"abc123",
		},
		Attachments: []string{
			"invoice.zip",
		},
		MITRE:     "T1566",
		RiskScore: 850,
		RiskLevel: "CRITICAL",
	}

	tmp, err := os.CreateTemp("", "ioc-*.json")
	if err != nil {
		t.Fatal(err)
	}

	tmp.Close()
	defer os.Remove(tmp.Name())

	err = ExportIOC(
		report,
		tmp.Name(),
	)

	if err != nil {
		t.Fatal(err)
	}

	data, err := os.ReadFile(
		tmp.Name(),
	)

	if err != nil {
		t.Fatal(err)
	}

	var loaded IOCReport

	err = json.Unmarshal(
		data,
		&loaded,
	)

	if err != nil {
		t.Fatal(err)
	}

	if loaded.Sender != report.Sender {
		t.Errorf(
			"expected %s got %s",
			report.Sender,
			loaded.Sender,
		)
	}

	if loaded.RiskScore != report.RiskScore {
		t.Errorf(
			"expected %d got %d",
			report.RiskScore,
			loaded.RiskScore,
		)
	}

	if loaded.RiskLevel != report.RiskLevel {
		t.Errorf(
			"expected %s got %s",
			report.RiskLevel,
			loaded.RiskLevel,
		)
	}

	if len(loaded.URLs) != 1 ||
		loaded.URLs[0] != "https://evil.com" {

		t.Error("URLs not exported correctly")
	}

	if len(loaded.Domains) != 1 ||
		loaded.Domains[0] != "evil.com" {

		t.Error("Domains not exported correctly")
	}

	if len(loaded.Hashes) != 1 ||
		loaded.Hashes[0] != "abc123" {

		t.Error("Hashes not exported correctly")
	}

	if len(loaded.Attachments) != 1 ||
		loaded.Attachments[0] != "invoice.zip" {

		t.Error("Attachments not exported correctly")
	}
}

func TestExportIOC_MarshalError(t *testing.T) {

	original := marshalJSON
	defer func() {
		marshalJSON = original
	}()

	marshalJSON = func(v any, prefix, indent string) ([]byte, error) {
		return nil, errors.New("marshal failed")
	}

	err := ExportIOC(
		IOCReport{},
		"ignored.json",
	)

	if err == nil {
		t.Fatal("expected marshal error")
	}

	if err.Error() != "marshal failed" {
		t.Fatalf("expected 'marshal failed', got %v", err)
	}
}
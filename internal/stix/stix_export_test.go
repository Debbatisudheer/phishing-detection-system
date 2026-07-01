package stix

import (
	"encoding/json"
	"errors"
	"os"
	"strings"
	"testing"
)

func TestExportURLIndicator(t *testing.T) {

	tmp, err := os.CreateTemp("", "stix-*.json")
	if err != nil {
		t.Fatal(err)
	}

	tmp.Close()
	defer os.Remove(tmp.Name())

	err = ExportURLIndicator(
		"https://evil.com",
		tmp.Name(),
	)

	if err != nil {
		t.Fatal(err)
	}

	data, err := os.ReadFile(tmp.Name())
	if err != nil {
		t.Fatal(err)
	}

	var indicator Indicator

	err = json.Unmarshal(data, &indicator)
	if err != nil {
		t.Fatal(err)
	}

	if indicator.Type != "indicator" {
		t.Errorf("expected indicator got %s", indicator.Type)
	}

	if indicator.SpecVersion != "2.1" {
		t.Errorf("expected 2.1 got %s", indicator.SpecVersion)
	}

	if indicator.Name != "Malicious URL" {
		t.Errorf("expected Malicious URL got %s", indicator.Name)
	}

	expectedPattern := "[url:value = 'https://evil.com']"

	if indicator.Pattern != expectedPattern {
		t.Errorf(
			"expected %s got %s",
			expectedPattern,
			indicator.Pattern,
		)
	}

	if indicator.PatternType != "stix" {
		t.Errorf("expected stix got %s", indicator.PatternType)
	}

	if !strings.HasPrefix(
		indicator.ID,
		"indicator--",
	) {
		t.Error("invalid indicator id")
	}

	if indicator.Created == "" {
		t.Error("created timestamp missing")
	}

	if indicator.Modified == "" {
		t.Error("modified timestamp missing")
	}
}

func TestExportURLIndicator_MarshalError(t *testing.T) {

	original := marshalJSON

	defer func() {
		marshalJSON = original
	}()

	marshalJSON = func(v any, prefix, indent string) ([]byte, error) {
		return nil, errors.New("marshal failed")
	}

	err := ExportURLIndicator(
		"https://evil.com",
		"ignored.json",
	)

	if err == nil {
		t.Fatal("expected marshal error")
	}

	if err.Error() != "marshal failed" {
		t.Fatalf("unexpected error %v", err)
	}
}
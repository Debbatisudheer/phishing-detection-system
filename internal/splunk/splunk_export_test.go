package splunk

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestExportEvent(t *testing.T) {

	before, _ := filepath.Glob("splunk_event_*.json")

	event := SplunkEvent{
		Timestamp: "2026-07-01T12:00:00Z",
		EventType: "PHISHING",
		Sender:    "alice@example.com",
		Subject:   "Urgent Login",
		RiskScore: 900,
		Decision:  "QUARANTINE",
		MITRE:     "T1566",
	}

	err := ExportEvent(event)
	if err != nil {
		t.Fatal(err)
	}

	after, _ := filepath.Glob("splunk_event_*.json")

	var newFile string

	for _, f := range after {

		found := false

		for _, old := range before {

			if old == f {
				found = true
				break
			}
		}

		if !found {
			newFile = f
			break
		}
	}

	if newFile == "" {

		t.Fatal(
			"no splunk export file created",
		)
	}

	defer os.Remove(newFile)

	data, err := os.ReadFile(newFile)
	if err != nil {
		t.Fatal(err)
	}

	var loaded SplunkEvent

	err = json.Unmarshal(data, &loaded)
	if err != nil {
		t.Fatal(err)
	}

	if loaded.Sender != event.Sender {
		t.Errorf(
			"expected %s got %s",
			event.Sender,
			loaded.Sender,
		)
	}

	if loaded.Subject != event.Subject {
		t.Errorf(
			"expected %s got %s",
			event.Subject,
			loaded.Subject,
		)
	}

	if loaded.RiskScore != event.RiskScore {
		t.Errorf(
			"expected %d got %d",
			event.RiskScore,
			loaded.RiskScore,
		)
	}

	if loaded.Decision != event.Decision {
		t.Errorf(
			"expected %s got %s",
			event.Decision,
			loaded.Decision,
		)
	}

	if loaded.MITRE != event.MITRE {
		t.Errorf(
			"expected %s got %s",
			event.MITRE,
			loaded.MITRE,
		)
	}

	if !strings.Contains(
		loaded.Timestamp,
		"2026",
	) {
		t.Error(
			"timestamp not exported correctly",
		)
	}
}

func TestExportEvent_MarshalError(t *testing.T) {

	original := marshalJSON

	defer func() {
		marshalJSON = original
	}()

	marshalJSON = func(v any, prefix, indent string) ([]byte, error) {
		return nil, errors.New("marshal failed")
	}

	err := ExportEvent(
		SplunkEvent{},
	)

	if err == nil {
		t.Fatal(
			"expected marshal error",
		)
	}

	if err.Error() != "marshal failed" {
		t.Fatalf(
			"unexpected error %v",
			err,
		)
	}
}
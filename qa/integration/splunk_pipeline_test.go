package integration

import (
	"os"
	"path/filepath"
	"testing"

	"phishing-platform/internal/splunk"
)

func TestSplunkPipeline(t *testing.T) {

	before, _ := filepath.Glob("splunk_event_*.json")

	event := splunk.SplunkEvent{
		Timestamp: "2026-07-02T12:00:00Z",
		EventType: "Phishing Detection",
		Sender:    "attacker@example.com",
		Subject:   "Verify Your PayPal Account",
		RiskScore: 670,
		Decision:  "QUARANTINE",
		MITRE:     "T1566 - Phishing",
	}

	err := splunk.ExportEvent(event)

	if err != nil {
		t.Fatal(err)
	}

	after, _ := filepath.Glob("splunk_event_*.json")

	if len(after) <= len(before) {
		t.Fatal("Splunk event file was not created")
	}

	latest := after[len(after)-1]

	info, err := os.Stat(latest)

	if err != nil {
		t.Fatal(err)
	}

	if info.Size() == 0 {
		t.Fatal("Splunk event file is empty")
	}

	defer os.Remove(latest)

	t.Log("========= SPLUNK PIPELINE =========")
	t.Log("File:", latest)
	t.Log("Size:", info.Size())
	t.Log("==================================")
}
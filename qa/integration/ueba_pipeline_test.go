package integration

import (
	"testing"
	"time"

	"phishing-platform/internal/ueba"
)

func TestUEBAPipeline(t *testing.T) {

	previous := ueba.LoginEvent{
		User:      "alice",
		Country:   "India",
		Timestamp: time.Now(),
	}

	current := ueba.LoginEvent{
		User:      "alice",
		Country:   "USA",
		Timestamp: time.Now().Add(1 * time.Hour),
	}

	findings := ueba.DetectImpossibleTravel(
		previous,
		current,
	)

	t.Log("========= UEBA PIPELINE =========")
	t.Log("Previous Country:", previous.Country)
	t.Log("Current Country :", current.Country)
	t.Log("Findings:", findings)
	t.Log("=================================")

	if len(findings) == 0 {
		t.Fatal("expected impossible travel detection")
	}

	if findings[0] != "Impossible travel detected" {
		t.Fatal("unexpected finding")
	}
}

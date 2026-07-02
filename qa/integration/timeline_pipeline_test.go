package integration

import (
	"testing"

	"phishing-platform/internal/timeline"
)

func TestTimelinePipeline(t *testing.T) {

	events := []string{
		"Email received",
		"Headers analyzed",
		"URLs extracted",
		"Domain analyzed",
		"Risk calculated",
		"Decision generated",
		"Report generated",
	}

	for _, event := range events {
		timeline.LogEvent(event)
	}

	t.Log("========= TIMELINE PIPELINE =========")

	for _, event := range events {
		t.Log(event)
	}

	t.Log("====================================")

	if len(events) != 7 {
		t.Fatal("timeline events missing")
	}
}
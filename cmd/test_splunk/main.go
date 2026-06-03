package main

import (
	"fmt"

	"phishing-platform/internal/splunk"
)

func main() {

	event := splunk.SplunkEvent{
		Timestamp: "2026-06-01 15:30:00",
		EventType: "phishing_detected",
		Sender:    "attacker@evil.com",
		Subject:   "Microsoft Security Alert",
		RiskScore: 770,
		Decision:  "QUARANTINE",
		MITRE:     "T1566 - Phishing",
	}

	err :=
		splunk.ExportEvent(
			event,
		)

	if err != nil {
		panic(err)
	}

	fmt.Println(
		"Splunk Event Exported",
	)
}
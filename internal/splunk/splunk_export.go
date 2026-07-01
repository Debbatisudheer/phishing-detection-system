package splunk

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

var marshalJSON = json.MarshalIndent

type SplunkEvent struct {
	Timestamp string `json:"timestamp"`
	EventType string `json:"event_type"`
	Sender    string `json:"sender"`
	Subject   string `json:"subject"`
	RiskScore int    `json:"risk_score"`
	Decision  string `json:"decision"`
	MITRE     string `json:"mitre"`
}

func ExportEvent(
	event SplunkEvent,
) error {

	data, err :=
	marshalJSON(
		event,
		"",
		"  ",
	)
	
	if err != nil {
		return err
	}

	filename :=
		fmt.Sprintf(
			"splunk_event_%d.json",
			time.Now().Unix(),
		)

	return os.WriteFile(
		filename,
		data,
		0644,
	)
}
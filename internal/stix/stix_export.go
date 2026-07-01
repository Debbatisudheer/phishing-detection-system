package stix

import (
	"encoding/json"
	"os"
	"time"
)

var marshalJSON = json.MarshalIndent

type Indicator struct {
	Type       string `json:"type"`
	SpecVersion string `json:"spec_version"`
	ID         string `json:"id"`
	Created    string `json:"created"`
	Modified   string `json:"modified"`
	Name       string `json:"name"`
	Pattern    string `json:"pattern"`
	PatternType string `json:"pattern_type"`
}

func ExportURLIndicator(
	url string,
	filename string,
) error {

	now :=
		time.Now().UTC().Format(
			time.RFC3339,
		)

	indicator := Indicator{
		Type:        "indicator",
		SpecVersion: "2.1",
		ID:          "indicator--" + now,
		Created:     now,
		Modified:    now,
		Name:        "Malicious URL",
		Pattern:     "[url:value = '" + url + "']",
		PatternType: "stix",
	}

	data, err :=
	marshalJSON(
		indicator,
		"",
		"  ",
	)

	if err != nil {
		return err
	}

	return os.WriteFile(
		filename,
		data,
		0644,
	)
}
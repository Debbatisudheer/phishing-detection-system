package ioc

import (
	"encoding/json"
	"os"
)

var marshalJSON = json.MarshalIndent

type IOCReport struct {
	Sender      string   `json:"sender"`
	URLs        []string `json:"urls"`
	Domains     []string `json:"domains"`
	Hashes      []string `json:"hashes"`
	Attachments []string `json:"attachments"`
	MITRE       string   `json:"mitre"`
	RiskScore   int      `json:"risk_score"`
	RiskLevel   string   `json:"risk_level"`
}

func ExportIOC(
	report IOCReport,
	filename string,
) error {

	data, err := marshalJSON(
		report,
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
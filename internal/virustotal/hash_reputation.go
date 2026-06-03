package virustotal

import (
	"encoding/json"
	"fmt"
)

type VTResponse struct {
	Data struct {
		Attributes struct {
			LastAnalysisStats struct {
				Malicious  int `json:"malicious"`
				Suspicious int `json:"suspicious"`
				Harmless   int `json:"harmless"`
			} `json:"last_analysis_stats"`
		} `json:"attributes"`
	} `json:"data"`
}

func CheckHashReputation(
	jsonData []byte,
) []string {

	var findings []string

	var response VTResponse

	err :=
		json.Unmarshal(
			jsonData,
			&response,
		)

	if err != nil {
		return findings
	}

	stats :=
		response.Data.Attributes.LastAnalysisStats

	fmt.Println(
	"VT Malicious:",
	stats.Malicious,
)

fmt.Println(
	"VT Suspicious:",
	stats.Suspicious,
)

fmt.Println(
	"VT Harmless:",
	stats.Harmless,
)

	if stats.Malicious > 0 {

		findings = append(
			findings,
			"VirusTotal malicious hash detected",
		)
	}

	if stats.Suspicious > 0 {

		findings = append(
			findings,
			"VirusTotal suspicious hash detected",
		)
	}

	return findings
}
package virustotal

import (
	"encoding/json"
)

type VTURLResponse struct {
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

func CheckURLReputation(
	jsonData []byte,
) []string {

	findings := []string{}

	var response VTURLResponse

	err := json.Unmarshal(
		jsonData,
		&response,
	)

	if err != nil {
		return findings
	}

	stats :=
		response.Data.Attributes.LastAnalysisStats

	if stats.Malicious > 0 {

	findings = append(
		findings,
		"VirusTotal malicious URL detected",
	)

} else if stats.Suspicious > 0 {

	findings = append(
		findings,
		"VirusTotal suspicious URL detected",
	)
}
	return findings
}
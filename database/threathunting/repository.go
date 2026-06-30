package threathunting

import (
    "phishing-platform/database"
    "strings"
)

func GetThreatHuntingStats() (
	map[string]interface{},
	error,
) {

	totalCritical := 0
	totalQuarantine := 0

	rows, err := database.DB.Query(
		`SELECT
			risk_level,
			verdict,
			mitre
		FROM analysis_results`,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	mitreMap :=
		make(map[string]int)

	for rows.Next() {

		var riskLevel string
		var verdict string
		var mitre string

		rows.Scan(
			&riskLevel,
			&verdict,
			&mitre,
		)

		if riskLevel == "CRITICAL" {
			totalCritical++
		}

		if verdict == "QUARANTINE" {
			totalQuarantine++
		}

		if mitre != "" {

	techniques :=
		strings.Split(
			mitre,
			"\n",
		)

	for _, technique := range techniques {

		technique =
			strings.TrimSpace(
				technique,
			)

		if technique == "" {
			continue
		}

		mitreMap[
			technique,
		]++
	}
}
	}

	return map[string]interface{}{
		"critical_files":   totalCritical,
		"quarantine_files": totalQuarantine,
		"top_mitre":        mitreMap,
	}, nil
}

package dashboard

import (
    "phishing-platform/database"
)

func GetDashboardStats() (
	int,
	int,
	int,
	int,
	int,
	error,
) {

	rows, err := database.DB.Query(
		`SELECT
			risk_level,
			verdict
		FROM analysis_results`,
	)

	if err != nil {
		return 0, 0, 0, 0, 0, err
	}

	defer rows.Close()

	total := 0
	allow := 0
	suspicious := 0
	quarantine := 0
	critical := 0

	for rows.Next() {

		var riskLevel string
		var verdict string

		rows.Scan(
			&riskLevel,
			&verdict,
		)

		total++

		switch verdict {

		case "ALLOW":
			allow++

		case "SUSPICIOUS":
			suspicious++

		case "QUARANTINE":
			quarantine++
		}

		if riskLevel == "CRITICAL" {
			critical++
		}
	}

	return total,
		allow,
		suspicious,
		quarantine,
		critical,
		nil
}

func GetRecentFindings() (
	[]map[string]interface{},
	error,
) {

	rows, err := database.DB.Query(
		`SELECT
			file_name,
			risk_score,
			risk_level,
			verdict
		FROM analysis_results
		ORDER BY id DESC
		LIMIT 10`,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	results := make(
	[]map[string]interface{},
	0,
)

	for rows.Next() {

		var fileName string
		var riskScore int
		var riskLevel string
		var verdict string

		rows.Scan(
			&fileName,
			&riskScore,
			&riskLevel,
			&verdict,
		)

		results = append(
			results,
			map[string]interface{}{
				"file_name":  fileName,
				"risk_score": riskScore,
				"risk_level": riskLevel,
				"verdict":    verdict,
			},
		)
	}

	return results, nil
}




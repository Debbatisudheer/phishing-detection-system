package search

import (
    "phishing-platform/database"
)

func SearchAnalysisResults(
	query string,
) ([]map[string]interface{}, error) {

	rows, err := database.DB.Query(
		`SELECT
			file_name,
			risk_score,
			risk_level,
			verdict,
			findings,
			sha256,
			urls,
			mitre
		FROM analysis_results
		WHERE
			file_name ILIKE '%' || $1 || '%'
			OR findings ILIKE '%' || $1 || '%'
			OR sha256 ILIKE '%' || $1 || '%'
			OR urls ILIKE '%' || $1 || '%'
			OR mitre ILIKE '%' || $1 || '%'
		ORDER BY id DESC`,
		query,
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
		var findings string
		var sha256 string
		var urls string
		var mitre string

		rows.Scan(
			&fileName,
			&riskScore,
			&riskLevel,
			&verdict,
			&findings,
			&sha256,
			&urls,
			&mitre,
		)

		results = append(
			results,
			map[string]interface{}{
				"file_name":  fileName,
				"risk_score": riskScore,
				"risk_level": riskLevel,
				"verdict":    verdict,
				"findings":   findings,
				"sha256":     sha256,
				"urls":       urls,
				"mitre":      mitre,
			},
		)
	}

	return results, nil
}

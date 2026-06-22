package database

func GetThreatIntelStats() (
	map[string]interface{},
	error,
) {

	var totalIOCs int
	var totalAlerts int
	var criticalFiles int

	DB.QueryRow(`
		SELECT COUNT(*)
		FROM ioc_correlation
	`).Scan(&totalIOCs)

	DB.QueryRow(`
		SELECT COUNT(*)
		FROM alerts
	`).Scan(&totalAlerts)

	DB.QueryRow(`
		SELECT COUNT(*)
		FROM sandbox_reports
		WHERE risk_level='CRITICAL'
	`).Scan(&criticalFiles)

	return map[string]interface{}{
		"total_iocs": totalIOCs,
		"total_alerts": totalAlerts,
		"critical_files": criticalFiles,
	}, nil
}

func GetTopRiskFiles() (
	[]map[string]interface{},
	error,
) {

	rows, err := DB.Query(`
	SELECT DISTINCT ON (file_name)
		file_name,
		risk_score,
		risk_level
	FROM sandbox_reports
	WHERE
		file_name IS NOT NULL
		AND file_name <> ''
	ORDER BY
		file_name,
		risk_score DESC
	LIMIT 10
`)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var results []map[string]interface{}

	for rows.Next() {

		var file string
		var score int
		var level string

		rows.Scan(
			&file,
			&score,
			&level,
		)

		results = append(
			results,
			map[string]interface{}{
				"file": file,
				"score": score,
				"level": level,
			},
		)
	}

	return results, nil
}
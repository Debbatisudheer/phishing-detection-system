package database

func SaveAnalysisResult(
	fileName string,
	riskScore int,
	riskLevel string,
	verdict string,
) error {

	_, err := DB.Exec(
		`INSERT INTO analysis_results
		(file_name, risk_score, risk_level, verdict)
		VALUES ($1, $2, $3, $4)`,
		fileName,
		riskScore,
		riskLevel,
		verdict,
	)

	return err
}

func GetAllAnalysisResults() (
	[]map[string]interface{},
	error,
) {

	rows, err := DB.Query(
		`SELECT
			file_name,
			risk_score,
			risk_level,
			verdict
		FROM analysis_results
		ORDER BY id DESC`,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var results []map[string]interface{}

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

func GetHighRiskAnalysisResults() (
	[]map[string]interface{},
	error,
) {

	rows, err := DB.Query(
		`SELECT
			file_name,
			risk_score,
			risk_level,
			verdict
		FROM analysis_results
		WHERE risk_level = 'HIGH'
		   OR risk_level = 'CRITICAL'
		ORDER BY id DESC`,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var results []map[string]interface{}

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

func GetDashboardStats() (
	int,
	int,
	int,
	int,
	int,
	error,
) {

	rows, err := DB.Query(
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
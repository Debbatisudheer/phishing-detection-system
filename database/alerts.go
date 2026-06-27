package database

func GetAlerts() (
	[]map[string]interface{},
	error,
) {

	rows, err := DB.Query(
		`
		SELECT
			id,
			alert_time,
			file_name,
			risk_level,
			verdict,
			message
		FROM alerts
		ORDER BY id DESC
		`,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	// IMPORTANT:
	// Initialize as an empty slice so JSON becomes [] instead of null.
	alerts := []map[string]interface{}{}

	for rows.Next() {

		var id int
		var alertTime string
		var fileName string
		var riskLevel string
		var verdict string
		var message string

		err := rows.Scan(
			&id,
			&alertTime,
			&fileName,
			&riskLevel,
			&verdict,
			&message,
		)

		if err != nil {
			return alerts, err
		}

		alerts = append(
			alerts,
			map[string]interface{}{
				"id":         id,
				"alert_time": alertTime,
				"file_name":  fileName,
				"risk_level": riskLevel,
				"verdict":    verdict,
				"message":    message,
			},
		)
	}

	return alerts, nil
}
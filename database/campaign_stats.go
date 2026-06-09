package database

func GetCampaignStats() (
	map[string]interface{},
	error,
) {

	rows, err := DB.Query(`
		SELECT
			risk_level
		FROM alerts
		WHERE verdict='CAMPAIGN'
	`)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	total := 0
	medium := 0
	high := 0
	critical := 0

	for rows.Next() {

		var risk string

		rows.Scan(&risk)

		total++

		if risk == "MEDIUM" {
			medium++
		}

		if risk == "HIGH" {
			high++
		}

		if risk == "CRITICAL" {
			critical++
		}
	}

	return map[string]interface{}{
		"total_campaigns": total,
		"medium": medium,
		"high": high,
		"critical": critical,
	}, nil
}
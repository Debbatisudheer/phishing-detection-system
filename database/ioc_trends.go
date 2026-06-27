package database

func GetIOCTrends() (
	[]map[string]interface{},
	error,
) {

	rows, err := DB.Query(`
		SELECT
			DATE(created_at) as trend_date,
			SUM(hit_count) as count
		FROM ioc_correlation
		GROUP BY DATE(created_at)
		ORDER BY trend_date
	`)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	results := make(
	[]map[string]interface{},
	0,
)

	for rows.Next() {

		var trendDate string
		var count int

		rows.Scan(
			&trendDate,
			&count,
		)

		results = append(
			results,
			map[string]interface{}{
				"date": trendDate,
				"count": count,
			},
		)
	}

	return results, nil
}
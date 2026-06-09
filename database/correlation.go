package database

func GetCorrelatedIOCs() (
	[]map[string]interface{},
	error,
) {

	rows, err := DB.Query(`
		SELECT
			ioc,
			COUNT(*) as count
		FROM ioc_correlation
		GROUP BY ioc
		HAVING COUNT(*) > 1
		ORDER BY count DESC
	`)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var results []map[string]interface{}

	for rows.Next() {

		var ioc string
		var count int

		rows.Scan(
			&ioc,
			&count,
		)

		results = append(
			results,
			map[string]interface{}{
				"ioc":   ioc,
				"count": count,
			},
		)
	}

	return results, nil
}
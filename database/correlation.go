package database

func GetCorrelatedIOCsDetailed() (
	[]map[string]interface{},
	error,
) {

	rows, err := DB.Query(`
		SELECT
			ioc,
			SUM(hit_count) as count,
			STRING_AGG(
				DISTINCT source_type,
				', '
			) as sources
		FROM ioc_correlation
		GROUP BY ioc
		HAVING COUNT(*) > 1
		ORDER BY count DESC
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

		var ioc string
		var count int
		var sources string

		rows.Scan(
			&ioc,
			&count,
			&sources,
		)

		results = append(
			results,
			map[string]interface{}{
				"ioc":     ioc,
				"count":   count,
				"sources": sources,
			},
		)
	}

	return results, nil
}
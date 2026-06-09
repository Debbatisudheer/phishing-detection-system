package database

func GetIOCSources(
	ioc string,
) (
	[]map[string]interface{},
	error,
) {

	rows, err := DB.Query(
		`
		SELECT
			source_type,
			file_name,
			created_at
		FROM ioc_correlation
		WHERE ioc = $1
		ORDER BY created_at DESC
		`,
		ioc,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var results []map[string]interface{}

	for rows.Next() {

		var sourceType string
		var fileName string
		var createdAt string

		rows.Scan(
			&sourceType,
			&fileName,
			&createdAt,
		)

		results = append(
			results,
			map[string]interface{}{
				"source_type": sourceType,
				"file_name": fileName,
				"created_at": createdAt,
			},
		)
	}

	return results, nil
}
package database

func GetIOCGraph() (
	[]map[string]interface{},
	error,
) {

	rows, err := DB.Query(`
		SELECT
			ioc,
			source_type,
			file_name
		FROM ioc_correlation
		ORDER BY ioc
	`)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var results []map[string]interface{}

	for rows.Next() {

		var ioc string
		var sourceType string
		var fileName string

		rows.Scan(
			&ioc,
			&sourceType,
			&fileName,
		)

		results = append(
			results,
			map[string]interface{}{
				"ioc": ioc,
				"source": sourceType,
				"file": fileName,
			},
		)
	}

	return results, nil
}
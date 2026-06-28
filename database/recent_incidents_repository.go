package database

func GetRecentIncidents() (
	[]map[string]interface{},
	error,
) {

	rows, err := DB.Query(
		`SELECT
			id,
			file_name,
			analyst,
			status
		FROM cases
		ORDER BY id DESC
		LIMIT 10`,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var incidents []map[string]interface{}
	

	for rows.Next() {

		var id int
		var fileName string
		var analyst string
		var status string

		rows.Scan(
			&id,
			&fileName,
			&analyst,
			&status,
		)

		incidents = append(
			incidents,
			map[string]interface{}{
				"id": id,
				"file_name": fileName,
				"analyst": analyst,
				"status": status,
			},
		)
	}

	return incidents, nil
}
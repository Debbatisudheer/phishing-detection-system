package database

func SaveAnalystNote(
	ioc string,
	analyst string,
	notes string,
) error {

	_, err := DB.Exec(`
		INSERT INTO analyst_notes (
			ioc,
			analyst,
			notes
		)
		VALUES ($1,$2,$3)
	`,
		ioc,
		analyst,
		notes,
	)

	return err
}

func GetAnalystNotes(
	ioc string,
) (
	[]map[string]interface{},
	error,
) {

	rows, err := DB.Query(`
		SELECT
			analyst,
			notes,
			created_at
		FROM analyst_notes
		WHERE ioc = $1
		ORDER BY created_at DESC
	`,
		ioc,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	results := make(
	[]map[string]interface{},
	0,
)

	for rows.Next() {

		var analyst string
		var notes string
		var created string

		rows.Scan(
			&analyst,
			&notes,
			&created,
		)

		results = append(
			results,
			map[string]interface{}{
				"analyst": analyst,
				"notes": notes,
				"created_at": created,
			},
		)
	}

	return results, nil
}
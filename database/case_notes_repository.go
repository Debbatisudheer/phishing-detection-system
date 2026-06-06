package database

func AddCaseNote(
	caseID int,
	analyst string,
	note string,
) error {

	_, err := DB.Exec(
		`INSERT INTO case_notes
		(case_id, analyst, note)
		VALUES ($1,$2,$3)`,
		caseID,
		analyst,
		note,
	)

	return err
}

func GetCaseNotes(
	caseID int,
) ([]map[string]interface{}, error) {

	rows, err := DB.Query(
		`SELECT
			analyst,
			note,
			created_at
		FROM case_notes
		WHERE case_id = $1
		ORDER BY created_at DESC`,
		caseID,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var notes []map[string]interface{}

	for rows.Next() {

		var analyst string
		var note string
		var createdAt string

		rows.Scan(
			&analyst,
			&note,
			&createdAt,
		)

		notes = append(
			notes,
			map[string]interface{}{
				"analyst": analyst,
				"note": note,
				"created_at": createdAt,
			},
		)
	}

	return notes, nil
}
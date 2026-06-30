package database

import (
    database "phishing-platform/database"
)

func CreateCase(
	fileName string,
	analyst string,
) error {

	_, err := database.DB.Exec(
		`INSERT INTO cases
		(file_name, analyst, status, notes)
		VALUES ($1, $2, $3, $4)`,
		fileName,
		analyst,
		"OPEN",
		"",
	)

	return err
}

func GetAllCases() (
	[]map[string]interface{},
	error,
) {

	rows, err := database.DB.Query(
		`SELECT
			id,
			file_name,
			analyst,
			status,
			notes,
			created_at
		FROM cases
		ORDER BY id DESC`,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	cases := []map[string]interface{}{}

	for rows.Next() {

		var id int
		var fileName string
		var analyst string
		var status string
		var notes string
		var createdAt string

		rows.Scan(
			&id,
			&fileName,
			&analyst,
			&status,
			&notes,
			&createdAt,
		)

		cases = append(
			cases,
			map[string]interface{}{
				"id":         id,
				"file_name":  fileName,
				"analyst":    analyst,
				"status":     status,
				"notes":      notes,
				"created_at": createdAt,
			},
		)
	}

	return cases, nil
}

func UpdateCase(
	id string,
	status string,
	notes string,
) error {

	_, err := database.DB.Exec(
		`UPDATE cases
		SET status = $1,
		    notes = $2
		WHERE id = $3`,
		status,
		notes,
		id,
	)

	return err
}

func GetCaseByID(
	id string,
) (map[string]interface{}, error) {

	var caseID int
	var fileName string
	var analyst string
	var status string
	var notes string
	var createdAt string

	err := database.DB.QueryRow(
		`SELECT
			id,
			file_name,
			analyst,
			status,
			notes,
			created_at
		FROM cases
		WHERE id = $1`,
		id,
	).Scan(
		&caseID,
		&fileName,
		&analyst,
		&status,
		&notes,
		&createdAt,
	)

	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"id":         caseID,
		"file_name":  fileName,
		"analyst":    analyst,
		"status":     status,
		"notes":      notes,
		"created_at": createdAt,
	}, nil
}

func CloseCase(
	id string,
) error {

	_, err := database.DB.Exec(
		`UPDATE cases
		SET status = 'CLOSED'
		WHERE id = $1`,
		id,
	)

	return err
}
package database

func GetIncidentStats() (
	map[string]int,
	error,
) {

	var open int
	var closed int

	DB.QueryRow(
		`SELECT COUNT(*)
		FROM cases
		WHERE status='OPEN'`,
	).Scan(
		&open,
	)

	DB.QueryRow(
		`SELECT COUNT(*)
		FROM cases
		WHERE status='CLOSED'`,
	).Scan(
		&closed,
	)

	return map[string]int{
		"open": open,
		"closed": closed,
		"total": open + closed,
	}, nil
}
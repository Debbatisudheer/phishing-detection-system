package database

func CreateUser(
	username string,
	password string,
) error {

	_, err := DB.Exec(
		`INSERT INTO users
		(username, password)
		VALUES ($1, $2)`,
		username,
		password,
	)

	return err
}

func GetUserByUsername(
	username string,
) (string, error) {

	var password string

	err := DB.QueryRow(
		`SELECT password
		FROM users
		WHERE username = $1`,
		username,
	).Scan(
		&password,
	)

	if err != nil {
		return "", err
	}

	return password, nil
}
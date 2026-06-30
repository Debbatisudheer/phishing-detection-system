package users

import (
    "phishing-platform/database"
)

func CreateUser(
	username string,
	password string,
) error {

	_, err := database.DB.Exec(
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
) (
	string,
	string,
	error,
) {

	var password string
	var role string

	err := database.DB.QueryRow(
		`SELECT password, role
		FROM users
		WHERE username = $1`,
		username,
	).Scan(
		&password,
		&role,
	)

	if err != nil {

		return "",
			"",
			err
	}

	return password,
		role,
		nil
}
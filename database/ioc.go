package database

func SaveIOC(
	ioc string,
	sourceType string,
	fileName string,
) error {

	var existing int

	err := DB.QueryRow(
		`
		SELECT COUNT(*)
		FROM ioc_correlation
		WHERE ioc=$1
		`,
		ioc,
	).Scan(&existing)

	if err != nil {
		return err
	}

	if existing > 0 {

		_, err = DB.Exec(
			`
			UPDATE ioc_correlation
			SET
				hit_count = hit_count + 1,
				last_seen = NOW()
			WHERE ioc=$1
			`,
			ioc,
		)

		return err
	}

	_, err = DB.Exec(
		`
		INSERT INTO ioc_correlation(
			ioc,
			source_type,
			file_name,
			first_seen,
			last_seen,
			hit_count
		)
		VALUES(
			$1,$2,$3,
			NOW(),
			NOW(),
			1
		)
		`,
		ioc,
		sourceType,
		fileName,
	)

	return err
}
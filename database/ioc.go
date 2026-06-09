package database

func SaveIOC(
	ioc string,
	sourceType string,
	fileName string,
) error {

	query := `
	INSERT INTO ioc_correlation (
		ioc,
		source_type,
		file_name
	)
	VALUES ($1,$2,$3)
	`

	_, err := DB.Exec(
		query,
		ioc,
		sourceType,
		fileName,
	)

	return err
}
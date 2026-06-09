package database

func SaveCampaignAlert(
	ioc string,
	count int,
	severity string,
) error {

	var existing int

	DB.QueryRow(
		`
		SELECT COUNT(*)
		FROM alerts
		WHERE campaign_ioc=$1
		`,
		ioc,
	).Scan(
		&existing,
	)

	if existing > 0 {

		return nil
	}

	message :=
		"CAMPAIGN DETECTED | IOC=" +
			ioc

	query := `
	INSERT INTO alerts (
		file_name,
		risk_level,
		verdict,
		message,
		campaign_ioc
	)
	VALUES ($1,$2,$3,$4,$5)
	`

	_, err := DB.Exec(
		query,
		ioc,
		severity,
		"CAMPAIGN",
		message,
		ioc,
	)

	return err
}
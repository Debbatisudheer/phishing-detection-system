package alerts

import "phishing-platform/database"

func SaveAlert(
	fileName string,
	riskLevel string,
	verdict string,
	message string,
) error {

	query := `
	INSERT INTO alerts (
		file_name,
		risk_level,
		verdict,
		message
	)
	VALUES ($1,$2,$3,$4)
	`

	_, err := database.DB.Exec(
		query,
		fileName,
		riskLevel,
		verdict,
		message,
	)

	return err
}


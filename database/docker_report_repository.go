package database

func SaveDockerReport(
	jobID int,
	containerStatus string,
	executionStatus string,
	duration int,
) error {

	query := `
	INSERT INTO docker_reports (
		sandbox_job_id,
		container_status,
		execution_status,
		duration_seconds
	)
	VALUES ($1,$2,$3,$4)
	`

	_, err := DB.Exec(
		query,
		jobID,
		containerStatus,
		executionStatus,
		duration,
	)

	return err
}
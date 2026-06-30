package database


import (
	"time"
	"fmt"
	 "phishing-platform/database"
)

type SandboxJob struct {
	ID          int       `json:"id"`
	FileName    *string    `json:"file_name"`
	FilePath    string    `json:"file_path"`
	Status      string    `json:"status"`
	SubmittedAt time.Time `json:"submitted_at"`
	CompletedAt *time.Time `json:"completed_at"`
}

func CreateSandboxJob(
	fileName string,
	filePath string,
) (
	int,
	error,
) {

	var jobID int

	err := database.DB.QueryRow(
		`
		INSERT INTO sandbox_jobs
		(
			file_name,
			file_path,
			status
		)
		VALUES
		(
			$1,
			$2,
			'PENDING'
		)
		RETURNING id
		`,
		fileName,
		filePath,
	).Scan(
		&jobID,
	)

	if err != nil {
		return 0, err
	}

	return jobID, nil
}

func GetSandboxJobs() (
	[]SandboxJob,
	error,
) {

	rows, err := database.DB.Query(
		`
		SELECT
			id,
			file_name,
			file_path,
			status,
			submitted_at,
			completed_at
		FROM sandbox_jobs
		ORDER BY id DESC
		`,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var jobs []SandboxJob

	for rows.Next() {

		var job SandboxJob

		err = rows.Scan(
	&job.ID,
	&job.FileName,
	&job.FilePath,
	&job.Status,
	&job.SubmittedAt,
	&job.CompletedAt,
)

if err != nil {

	fmt.Println(
		"SCAN ERROR:",
		err,
	)

	continue
}

		jobs = append(
			jobs,
			job,
		)
	}

	return jobs, nil
}

func UpdateSandboxJobStatus(
	id int,
	status string,
) error {

	_, err := database.DB.Exec(
		`
		UPDATE sandbox_jobs
		SET
			status = $1,
			completed_at = NOW()
		WHERE id = $2
		`,
		status,
		id,
	)

	return err
}
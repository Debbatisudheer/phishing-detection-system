package database

import (
	"time"
	"fmt"
	"phishing-platform/internal/models"
)

type SandboxReport struct {
	ID        int       `json:"id"`
	JobID     int       `json:"job_id"`

	FileName  string    `json:"file_name"`
	FileSize  int64     `json:"file_size"`
	Extension string    `json:"extension"`
	MIMEType  string    `json:"mime_type"`
	MD5       string    `json:"md5"`
	SHA256    string    `json:"sha256"`

	Findings  string    `json:"findings"`
	RiskScore int       `json:"risk_score"`
	RiskLevel string    `json:"risk_level"`
	Verdict   string    `json:"verdict"`
	Mitre     string    `json:"mitre"`
	CreatedAt time.Time `json:"created_at"`
}

func SaveSandboxReport(
	jobID int,
	fileName string,
	fileSize int64,
	extension string,
	mimeType string,
	md5 string,
	sha256 string,
	findings string,
	riskScore int,
	riskLevel string,
	verdict string,
	mitre string,
) error {

	result, err := DB.Exec(
		`
		INSERT INTO sandbox_reports
		(
			job_id,
			file_name,
			file_size,
			extension,
			mime_type,
			md5,
			sha256,
			findings,
			risk_score,
			risk_level,
			verdict,
			mitre
		)
		VALUES
		(
			$1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12
		)
		`,
		jobID,
		fileName,
		fileSize,
		extension,
		mimeType,
		md5,
		sha256,
		findings,
		riskScore,
		riskLevel,
		verdict,
		mitre,
	)

	if err != nil {

    fmt.Println(
        "DATABASE INSERT ERROR:",
        err,
    )

    return err
}

rows, _ := result.RowsAffected()

fmt.Println(
    "ROWS INSERTED:",
    rows,
)

return nil
}

func GetSandboxReportByID(
    id int,
) (
    models.SandboxReport,
    error,
) {

    var report models.SandboxReport

    err := DB.QueryRow(`
        SELECT
            id,
            file_name,
            file_size,
            extension,
            mime_type,
            md5,
            sha256,
            findings,
            risk_score,
            risk_level,
            verdict,
            mitre
        FROM sandbox_reports
        WHERE id=$1
    `, id).Scan(
        &report.ID,
        &report.FileName,
        &report.FileSize,
        &report.Extension,
        &report.MimeType,
        &report.MD5,
        &report.SHA256,
        &report.Findings,
        &report.RiskScore,
        &report.RiskLevel,
        &report.Verdict,
        &report.MITRE,
    )

    return report, err
}

func GetSandboxReportByJobID(
    jobID int,
) (
    models.SandboxReport,
    error,
) {

    var report models.SandboxReport

    err := DB.QueryRow(`
        SELECT
            id,
            job_id,
            file_name,
            file_size,
            extension,
            mime_type,
            md5,
            sha256,
            findings,
            risk_score,
            risk_level,
            verdict,
            mitre
        FROM sandbox_reports
        WHERE job_id = $1
    `, jobID).Scan(
        &report.ID,
        &report.JobID,
        &report.FileName,
        &report.FileSize,
        &report.Extension,
        &report.MimeType,
        &report.MD5,
        &report.SHA256,
        &report.Findings,
        &report.RiskScore,
        &report.RiskLevel,
        &report.Verdict,
        &report.MITRE,
    )

    return report, err
}
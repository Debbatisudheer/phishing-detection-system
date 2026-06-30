package database

import (
	"phishing-platform/internal/models"
	rootdb "phishing-platform/database"
)

func GetSandboxReports() (
	[]models.SandboxReport,
	error,
) {

	rows, err := rootdb.DB.Query(`
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
		ORDER BY id DESC
	`)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var reports []models.SandboxReport

	for rows.Next() {

		var report models.SandboxReport

		err := rows.Scan(
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

		if err != nil {
			return nil, err
		}

		reports = append(
			reports,
			report,
		)
	}

	return reports, nil
}
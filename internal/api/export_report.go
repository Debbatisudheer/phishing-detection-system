package api

import (
	"net/http"

	"phishing-platform/database"
	"phishing-platform/internal/report"
)

func ExportReportHandler(
	w http.ResponseWriter,
	r *http.Request,
) {

	var fileName string
	var riskScore int
	var riskLevel string
	var verdict string
	var findings string
	var mitre string

	err :=
		database.DB.QueryRow(`
			SELECT
				file_name,
				risk_score,
				risk_level,
				verdict,
				findings,
				mitre
			FROM sandbox_reports
			ORDER BY id DESC
			LIMIT 1
		`).Scan(
			&fileName,
			&riskScore,
			&riskLevel,
			&verdict,
			&findings,
			&mitre,
		)

	if err != nil {

		http.Error(
			w,
			err.Error(),
			http.StatusInternalServerError,
		)

		return
	}

	err =
		report.GeneratePDFReport(
			fileName,
			riskScore,
			riskLevel,
			verdict,
			findings,
			mitre,
		)

	if err != nil {

		http.Error(
			w,
			err.Error(),
			http.StatusInternalServerError,
		)

		return
	}

	w.Write(
		[]byte(
			"PDF Report Generated Successfully",
		),
	)
}
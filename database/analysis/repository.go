package analysis

import (
    "phishing-platform/database"
)

func SaveAnalysisResult(
	fileName string,
	riskScore int,
	riskLevel string,
	verdict string,
	findings string,
	sha256 string,
	urls string,
	mitre string,
) error {

	_, err := database.DB.Exec(
		`INSERT INTO analysis_results
(file_name, risk_score, risk_level, verdict, findings, sha256, urls, mitre)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`,
		fileName,
		riskScore,
		riskLevel,
		verdict,
		findings,
		sha256,
		urls,
		mitre,
	)

	return err
}

func GetAllAnalysisResults() (
	[]map[string]interface{},
	error,
) {

	rows, err := database.DB.Query(
		`SELECT
			file_name,
			risk_score,
			risk_level,
			verdict,
			findings,
			sha256,
	urls,
	mitre
		FROM analysis_results
		ORDER BY id DESC`,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	results := make(
	[]map[string]interface{},
	0,
)

	for rows.Next() {

		var fileName string
		var riskScore int
		var riskLevel string
		var verdict string
		var findings string
		var sha256 string
var urls string
var mitre string

		rows.Scan(
			&fileName,
			&riskScore,
			&riskLevel,
			&verdict,
			&findings,
			&sha256,
&urls,
&mitre,
		)

		results = append(
			results,
			map[string]interface{}{
				"file_name":  fileName,
				"risk_score": riskScore,
				"risk_level": riskLevel,
				"verdict":    verdict,
				"findings":   findings,
				"sha256": sha256,
"urls": urls,
"mitre": mitre,
			},
		)
	}

	return results, nil
}

func GetHighRiskAnalysisResults() (
	[]map[string]interface{},
	error,
) {

	rows, err := database.DB.Query(
		`SELECT
			file_name,
			risk_score,
			risk_level,
			verdict,
			findings,
			sha256,
	urls,
	mitre
		FROM analysis_results
		WHERE risk_level = 'HIGH'
		   OR risk_level = 'CRITICAL'
		ORDER BY id DESC`,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	results := make(
	[]map[string]interface{},
	0,
)

	for rows.Next() {

		var fileName string
		var riskScore int
		var riskLevel string
		var verdict string
		var findings string
		var sha256 string
var urls string
var mitre string

		rows.Scan(
			&fileName,
			&riskScore,
			&riskLevel,
			&verdict,
			&findings,
			&sha256,
&urls,
&mitre,
		)

		results = append(
			results,
			map[string]interface{}{
				"file_name":  fileName,
				"risk_score": riskScore,
				"risk_level": riskLevel,
				"verdict":    verdict,
				"findings":   findings,
				"sha256": sha256,
"urls": urls,
"mitre": mitre,
			},
		)
	}

	return results, nil
}

func GetAnalysisResultByFileName(
	fileName string,
) (map[string]interface{}, error) {

	var riskScore int
var riskLevel string
var verdict string
var findings string
var sha256 string
var urls string
var mitre string

	err := database.DB.QueryRow(
		`SELECT
	risk_score,
	risk_level,
	verdict,
	findings,
	sha256,
	urls,
	mitre
		FROM analysis_results
		WHERE file_name = $1
		ORDER BY id DESC
		LIMIT 1`,
		fileName,
	).Scan(
		&riskScore,
&riskLevel,
&verdict,
&findings,
&sha256,
&urls,
&mitre,
	)

	if err != nil {
		return nil, err
	}

	result :=
		map[string]interface{}{
			"file_name":  fileName,
			"risk_score": riskScore,
			"risk_level": riskLevel,
			"verdict":    verdict,
			"findings":   findings,
			"sha256": sha256,
"urls": urls,
"mitre": mitre,
		}

	return result, nil
}

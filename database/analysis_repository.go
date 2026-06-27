package database

import (
	"strings"
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

	_, err := DB.Exec(
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

	rows, err := DB.Query(
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

	rows, err := DB.Query(
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

func GetDashboardStats() (
	int,
	int,
	int,
	int,
	int,
	error,
) {

	rows, err := DB.Query(
		`SELECT
			risk_level,
			verdict
		FROM analysis_results`,
	)

	if err != nil {
		return 0, 0, 0, 0, 0, err
	}

	defer rows.Close()

	total := 0
	allow := 0
	suspicious := 0
	quarantine := 0
	critical := 0

	for rows.Next() {

		var riskLevel string
		var verdict string

		rows.Scan(
			&riskLevel,
			&verdict,
		)

		total++

		switch verdict {

		case "ALLOW":
			allow++

		case "SUSPICIOUS":
			suspicious++

		case "QUARANTINE":
			quarantine++
		}

		if riskLevel == "CRITICAL" {
			critical++
		}
	}

	return total,
		allow,
		suspicious,
		quarantine,
		critical,
		nil
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

	err := DB.QueryRow(
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

func SearchAnalysisResults(
	query string,
) ([]map[string]interface{}, error) {

	rows, err := DB.Query(
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
		WHERE
			file_name ILIKE '%' || $1 || '%'
			OR findings ILIKE '%' || $1 || '%'
			OR sha256 ILIKE '%' || $1 || '%'
			OR urls ILIKE '%' || $1 || '%'
			OR mitre ILIKE '%' || $1 || '%'
		ORDER BY id DESC`,
		query,
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
				"sha256":     sha256,
				"urls":       urls,
				"mitre":      mitre,
			},
		)
	}

	return results, nil
}

func ExportIOCs() (
	map[string]interface{},
	error,
) {

	rows, err := DB.Query(
		`SELECT
			sha256,
			urls,
			mitre
		FROM analysis_results`,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var hashes []string
	var urls []string
	var mitres []string

	for rows.Next() {

		var sha256 string
		var url string
		var mitre string

		rows.Scan(
			&sha256,
			&url,
			&mitre,
		)

		if sha256 != "" {
			hashes = append(
				hashes,
				sha256,
			)
		}

		if url != "" {
			urls = append(
				urls,
				url,
			)
		}

		if mitre != "" {
			mitres = append(
				mitres,
				mitre,
			)
		}
	}

	return map[string]interface{}{
		"hashes": hashes,
		"urls":   urls,
		"mitre":  mitres,
	}, nil
}

func GetRecentFindings() (
	[]map[string]interface{},
	error,
) {

	rows, err := DB.Query(
		`SELECT
			file_name,
			risk_score,
			risk_level,
			verdict
		FROM analysis_results
		ORDER BY id DESC
		LIMIT 10`,
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

		rows.Scan(
			&fileName,
			&riskScore,
			&riskLevel,
			&verdict,
		)

		results = append(
			results,
			map[string]interface{}{
				"file_name":  fileName,
				"risk_score": riskScore,
				"risk_level": riskLevel,
				"verdict":    verdict,
			},
		)
	}

	return results, nil
}

func GetThreatHuntingStats() (
	map[string]interface{},
	error,
) {

	totalCritical := 0
	totalQuarantine := 0

	rows, err := DB.Query(
		`SELECT
			risk_level,
			verdict,
			mitre
		FROM analysis_results`,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	mitreMap :=
		make(map[string]int)

	for rows.Next() {

		var riskLevel string
		var verdict string
		var mitre string

		rows.Scan(
			&riskLevel,
			&verdict,
			&mitre,
		)

		if riskLevel == "CRITICAL" {
			totalCritical++
		}

		if verdict == "QUARANTINE" {
			totalQuarantine++
		}

		if mitre != "" {

	techniques :=
		strings.Split(
			mitre,
			"\n",
		)

	for _, technique := range techniques {

		technique =
			strings.TrimSpace(
				technique,
			)

		if technique == "" {
			continue
		}

		mitreMap[
			technique,
		]++
	}
}
	}

	return map[string]interface{}{
		"critical_files":   totalCritical,
		"quarantine_files": totalQuarantine,
		"top_mitre":        mitreMap,
	}, nil
}

func GetMITREStats() (
	[]map[string]interface{},
	error,
) {

	rows, err :=
		DB.Query(
			`SELECT mitre
			 FROM analysis_results`,
		)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	techniqueCount :=
		make(map[string]int)

	for rows.Next() {

		var mitre string

		rows.Scan(
			&mitre,
		)

		techniques :=
			strings.Split(
				mitre,
				"\n",
			)

		for _, technique := range techniques {

			technique =
				strings.TrimSpace(
					technique,
				)

			if technique == "" {
				continue
			}

			techniqueCount[
				technique,
			]++
		}
	}

	results := make(
	[]map[string]interface{},
	0,
)

	for technique, count :=
		range techniqueCount {

		results = append(
			results,
			map[string]interface{}{
				"technique": technique,
				"count": count,
			},
		)
	}

	return results, nil
}

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

	_, err := DB.Exec(
		query,
		fileName,
		riskLevel,
		verdict,
		message,
	)

	return err
}


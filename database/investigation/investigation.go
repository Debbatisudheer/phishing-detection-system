package database

import (
    "phishing-platform/database"
)

func GetInvestigationSummary(
	ioc string,
) (
	map[string]interface{},
	error,
) {

	rows, err := database.DB.Query(`
		SELECT
			source_type,
			file_name,
			created_at
		FROM ioc_correlation
		WHERE ioc = $1
		ORDER BY created_at
	`,
		ioc,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	sourceMap :=
		make(map[string]bool)

	fileMap :=
		make(map[string]bool)

	var sources []string
	var files []string

	count := 0

	var firstSeen string
	var lastSeen string

	for rows.Next() {

		var source string
		var file string
		var created string

		rows.Scan(
			&source,
			&file,
			&created,
		)

		if !sourceMap[source] {

			sourceMap[source] = true

			sources = append(
				sources,
				source,
			)
		}

		if !fileMap[file] {

			fileMap[file] = true

			files = append(
				files,
				file,
			)
		}

		if count == 0 {

			firstSeen = created
		}

		lastSeen = created

		count++
	}

	var riskLevel string
	var verdict string
	var mitre string

	database.DB.QueryRow(`
	SELECT
		risk_level,
		verdict,
		mitre
	FROM analysis_results
	ORDER BY id DESC
	LIMIT 1
`).Scan(
	&riskLevel,
	&verdict,
	&mitre,
)

	return map[string]interface{}{
		"ioc":         ioc,
		"count":       count,
		"sources":     sources,
		"files":       files,
		"first_seen":  firstSeen,
		"last_seen":   lastSeen,
		"risk_level":  riskLevel,
		"verdict":     verdict,
		"mitre":       mitre,
	}, nil
}
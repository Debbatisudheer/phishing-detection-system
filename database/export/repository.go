package export

import (
    "phishing-platform/database"
)


func ExportIOCs() (
	map[string]interface{},
	error,
) {

	rows, err := database.DB.Query(
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

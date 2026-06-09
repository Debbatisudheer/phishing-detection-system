package database

import "strings"

func GetMITREHeatmap() (
	[]map[string]interface{},
	error,
) {

	rows, err := DB.Query(`
		SELECT mitre
		FROM analysis_results
		WHERE mitre IS NOT NULL
		AND mitre <> ''
		AND mitre <> 'NO_MITRE_MATCH'
	`)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	techniqueCounts :=
		make(map[string]int)

	for rows.Next() {

		var mitre string

		rows.Scan(
			&mitre,
		)

		techniques :=
			strings.Split(
				mitre,
				",",
			)

		for _, technique :=
			range techniques {

			technique =
				strings.TrimSpace(
					technique,
				)

			if technique != "" {

				techniqueCounts[
					technique,
				]++
			}
		}
	}

	var results []map[string]interface{}

	for technique, count :=
		range techniqueCounts {

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
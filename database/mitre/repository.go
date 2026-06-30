package mitre

import (
    "phishing-platform/database"
    "strings"
)

func GetMITREStats() (
	[]map[string]interface{},
	error,
) {

	rows, err :=
		database.DB.Query(
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

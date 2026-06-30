package database

import (
    rootdb "phishing-platform/database"
)

func GetCampaignTimeline() (
	[]map[string]interface{},
	error,
) {

	rows, err := rootdb.DB.Query(`
		SELECT
			ioc,
			MIN(created_at) as first_seen,
			MAX(created_at) as last_seen,
			COUNT(*) as occurrences
		FROM ioc_correlation
		GROUP BY ioc
		ORDER BY occurrences DESC
	`)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	results := make(
	[]map[string]interface{},
	0,
)

	for rows.Next() {

		var ioc string
		var firstSeen string
		var lastSeen string
		var count int

		rows.Scan(
			&ioc,
			&firstSeen,
			&lastSeen,
			&count,
		)

		results = append(
			results,
			map[string]interface{}{
				"ioc":         ioc,
				"first_seen":  firstSeen,
				"last_seen":   lastSeen,
				"occurrences": count,
			},
		)
	}

	return results, nil
}
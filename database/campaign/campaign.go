package database

import (

	"fmt"
    rootdb "phishing-platform/database"
)

func GetCampaigns() (
	[]map[string]interface{},
	error,
) {

	rows, err := rootdb.DB.Query(`
		SELECT
			ioc,
			SUM(hit_count) as count,
			STRING_AGG(
				DISTINCT source_type,
				', '
			) as sources
		FROM ioc_correlation
		GROUP BY ioc
		HAVING SUM(hit_count) >= 2
		ORDER BY count DESC
	`)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var campaigns []map[string]interface{}

	for rows.Next() {

		var ioc string
		var count int
		var sources string

		err := rows.Scan(
			&ioc,
			&count,
			&sources,
		)

		if err != nil {
			continue
		}

		severity := "MEDIUM"

		if count >= 3 {
			severity = "HIGH"
		}

		if count >= 5 {
			severity = "CRITICAL"
		}

		if count >= 2 {

			err :=
				SaveCampaignAlert(
					ioc,
					count,
					severity,
				)

			if err != nil {

				fmt.Println(
					"Campaign Alert Error:",
					err,
				)
			}
		}

		campaigns = append(
			campaigns,
			map[string]interface{}{
				"ioc":      ioc,
				"count":    count,
				"severity": severity,
				"sources":  sources,
			},
		)
	}

	return campaigns, nil
}
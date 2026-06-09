package database

import "fmt"

func GetCampaigns() (
	[]map[string]interface{},
	error,
) {

	rows, err := DB.Query(`
		SELECT
			ioc,
			COUNT(*) as count
		FROM ioc_correlation
		GROUP BY ioc
		HAVING COUNT(*) >= 2
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

	rows.Scan(
		&ioc,
		&count,
	)

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
		},
	)
}

	return campaigns, nil
}
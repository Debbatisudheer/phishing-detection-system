package database

import (
    "phishing-platform/database"
)

func GetIOCReputation(
	ioc string,
) (
	map[string]interface{},
	error,
) {

	var reputation string
	var source string

	err := database.DB.QueryRow(`
		SELECT
			reputation,
			source
		FROM threat_intel_cache
		WHERE ioc = $1
	`,
		ioc,
	).Scan(
		&reputation,
		&source,
	)

	if err != nil {

		return map[string]interface{}{
			"reputation": "UNKNOWN",
			"source":     "N/A",
		}, nil
	}

	return map[string]interface{}{
		"reputation": reputation,
		"source":     source,
	}, nil
}
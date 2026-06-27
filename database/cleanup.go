package database

import (
	"fmt"
	"time"
)

var LastCleanupTime time.Time

func StartCleanupWorker() {

	go func() {

		ticker := time.NewTicker(
			30 * time.Minute, // Production
			//30 * time.Second, // Testing
		)

		defer ticker.Stop()

		for {

			<-ticker.C

			fmt.Println()
			fmt.Println("===================================")
			fmt.Println("Starting Database Cleanup...")
			fmt.Println("===================================")

			CleanupDatabase()

		}

	}()

}

func CleanupDatabase() {

	// -----------------------------------
	// Check whether there is any data
	// -----------------------------------

	var total int

	err := DB.QueryRow(`
		SELECT
			(SELECT COUNT(*) FROM analysis_results) +
			(SELECT COUNT(*) FROM sandbox_reports) +
			(SELECT COUNT(*) FROM sandbox_jobs) +
			(SELECT COUNT(*) FROM alerts) +
			(SELECT COUNT(*) FROM docker_reports)
	`).Scan(&total)

	if err != nil {

		fmt.Println(
			"Cleanup Check Error:",
			err,
		)

		return
	}

	if total == 0 {

		fmt.Println(
			"No records found. Cleanup skipped.",
		)

		fmt.Println("===================================")
		fmt.Println()

		return
	}

	// -----------------------------------
	// Cleanup
	// -----------------------------------

	tables := []string{

		"sandbox_reports",

		"sandbox_jobs",

		"analysis_results",

		"alerts",

		"docker_reports",

		// Add "iocs" here ONLY if that table exists.
	}

	for _, table := range tables {

		query := fmt.Sprintf(
			"TRUNCATE TABLE %s RESTART IDENTITY CASCADE",
			table,
		)

		_, err := DB.Exec(query)

		if err != nil {

			fmt.Println(
				"Cleanup Error:",
				table,
				err,
			)

			continue
		}

		fmt.Println(
			"Cleaned:",
			table,
		)
	}

	LastCleanupTime = time.Now()

	fmt.Println("===================================")
	fmt.Println("Database Cleanup Complete")
	fmt.Println(
		"Last Cleanup:",
		LastCleanupTime.Format("03:04:05 PM"),
	)
	fmt.Println("===================================")
	fmt.Println()

}
package ueba

import (
	"time"
)

type LoginEvent struct {
	User      string
	Country   string
	Timestamp time.Time
}

func DetectImpossibleTravel(
	previous LoginEvent,
	current LoginEvent,
) []string {

	var findings []string

	timeDifference :=
		current.Timestamp.Sub(
			previous.Timestamp,
		)

	// Simulated impossible travel
	if previous.Country != current.Country &&
		timeDifference.Hours() < 2 {

		findings = append(
			findings,
			"Impossible travel detected",
		)
	}

	return findings
}
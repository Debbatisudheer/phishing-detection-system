package sandbox

import "fmt"

func BuildTimeline(
	findings []string,
) []string {

	var timeline []string

	for i, finding := range findings {

		event :=
			fmt.Sprintf(
				"[00:%02d] %s",
				i+1,
				finding,
			)

		timeline = append(
			timeline,
			event,
		)
	}

	return timeline
}
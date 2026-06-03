package header

import (
	"strings"
)

func DetectDisplayNameSpoofing(
	from string,
) []string {

	var findings []string

	lower :=
		strings.ToLower(from)

	brands := []string{
		"microsoft",
		"paypal",
		"amazon",
		"google",
		"apple",
		"bank",
	}

	for _, brand := range brands {

		if strings.Contains(
			lower,
			brand,
		) &&
			!strings.Contains(
				lower,
				"@"+brand,
			) {

			findings = append(
				findings,
				"Display name spoofing detected: "+brand,
			)
		}
	}

	return findings
}
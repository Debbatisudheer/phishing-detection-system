package domain

import (
	"strings"
)

func DetectHomographDomain(
	host string,
) []string {

	var findings []string

	host = strings.ToLower(
		host,
	)


	suspiciousChars := []string{
		"а", // Cyrillic a
		"е", // Cyrillic e
		"о", // Cyrillic o
		"р", // Cyrillic p
		"с", // Cyrillic c
		"х", // Cyrillic x
		"і", // Cyrillic i
	}

	for _, ch := range suspiciousChars {

		if strings.Contains(
			host,
			ch,
		) {

			findings = append(
				findings,
				"Homograph domain detected",
			)

			break
		}
	}

	return findings
}
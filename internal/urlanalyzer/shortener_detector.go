package urlanalyzer

import "strings"

func DetectShortenedURL(
	url string,
) []string {

	var findings []string

	shorteners := []string{
		"bit.ly",
		"tinyurl.com",
		"t.co",
		"goo.gl",
		"is.gd",
	}

	for _, s := range shorteners {

		if strings.Contains(
			strings.ToLower(url),
			s,
		) {

			findings = append(
				findings,
				"Shortened URL detected: "+s,
			)
		}
	}

	return findings
}
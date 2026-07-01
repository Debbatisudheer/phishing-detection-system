package threatfeed

import (
	"strings"
)

var MaliciousDomains =
	LoadFeed(
		"feeds/openphish_feed.txt",
	)

func CheckThreatFeed(
	url string,
) []string {

	findings := []string{}

	url = strings.ToLower(
		url,
	)

	for _, domain := range MaliciousDomains {

		if strings.Contains(
			url,
			domain,
		) {

			findings = append(
				findings,
				"Threat feed hit: "+domain,
			)
		}
	}

	return findings
}
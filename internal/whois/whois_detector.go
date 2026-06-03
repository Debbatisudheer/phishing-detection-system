package whois

import (
	"strings"
)

func AnalyzeDomain(
	domain string,
) []string {

	var findings []string

	suspiciousTLDs := []string{
		".xyz",
		".top",
		".click",
		".shop",
		".online",
		".site",
	}

	for _, tld := range suspiciousTLDs {

		if strings.HasSuffix(
			strings.ToLower(domain),
			tld,
		) {

			findings = append(
				findings,
				"WHOIS suspicious TLD detected: "+tld,
			)
		}
	}

	return findings
}
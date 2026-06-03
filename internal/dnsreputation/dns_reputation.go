package dnsreputation

import (
	"strings"
)

func CheckDNSReputation(
	domain string,
) []string {

	var findings []string

	knownBadDomains := []string{
		"evil.com",
		"malware.com",
		"phishing-site.xyz",
	}

	for _, bad := range knownBadDomains {

		if strings.Contains(
			strings.ToLower(domain),
			bad,
		) {

			findings = append(
				findings,
				"DNS reputation hit: "+bad,
			)
		}
	}

	return findings
}
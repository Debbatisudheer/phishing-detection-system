package threatintel

import (
	"net/url"
)

var MaliciousDomains = []string{
	"evil.com",
	"paypal-security-login.xyz",
	"malicious-login.top",
}

var MaliciousIPs = []string{
	"185.99.22.10",
	"192.168.100.50",
}

func CheckThreatIntel(
	rawURL string,
) []string {

	var findings []string

	parsedURL, err := url.Parse(rawURL)

	if err != nil {
		return findings
	}

	host := parsedURL.Hostname()

	// Domain reputation
	for _, domain := range MaliciousDomains {

		if host == domain {

			findings = append(
				findings,
				"Known malicious domain detected: "+domain,
			)
		}
	}

	// IP reputation
	for _, ip := range MaliciousIPs {

		if host == ip {

			findings = append(
				findings,
				"Known malicious IP detected: "+ip,
			)
		}
	}

	return findings
}
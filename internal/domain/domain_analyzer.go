package domain

import (
	"net/url"
	"regexp"
	"strings"
	"fmt"
	"phishing-platform/internal/whois"
	"phishing-platform/internal/dnsreputation"
)

var SuspiciousTLDs = []string{
	".xyz",
	".top",
	".ru",
	".tk",
	".cn",
}

var SuspiciousBrands = []string{
	"paypal",
	"microsoft",
	"google",
	"apple",
	"amazon",
}

func AnalyzeURL(rawURL string) []string {

	var findings []string

	redirectFindings :=
	CheckRedirectURL(
		rawURL,
	)

findings = append(
	findings,
	redirectFindings...,
)

	parsedURL, err := url.Parse(rawURL)

	if err != nil {
		return findings
	}

	host := strings.ToLower(
		parsedURL.Host,
	)

	whoisFindings :=
	whois.AnalyzeDomain(
		host,
	)

findings = append(
	findings,
	whoisFindings...,
)

dnsFindings :=
	dnsreputation.CheckDNSReputation(
		host,
	)

findings = append(
		findings,
		dnsFindings...,
)
	domainAgeFindings :=
	CheckDomainAge(
		host,
	)

findings = append(
	findings,
	domainAgeFindings...,
)

	// Remove port if present
	host = strings.Split(
		host,
		":",
	)[0]

	homographFindings :=
	DetectHomographDomain(
		host,
	)

findings = append(
	findings,
	homographFindings...,
)

	// Lookalike domain detection
	lookalikeFindings :=
		DetectLookalikeDomain(
			host,
		)

	findings = append(
		findings,
		lookalikeFindings...,
	)

	// Suspicious TLD detection
	for _, tld := range SuspiciousTLDs {

		if strings.HasSuffix(
			host,
			tld,
		) {

			findings = append(
				findings,
				"Suspicious TLD detected: "+tld,
			)
		}
	}

	// IP URL detection
	ipRegex := regexp.MustCompile(
		`^(\d{1,3}\.){3}\d{1,3}$`,
	)

	if ipRegex.MatchString(host) {

		findings = append(
			findings,
			"IP-based URL detected",
		)
	}

	// Simulated suspicious country detection
	if strings.Contains(
		host,
		".ru",
	) {

		findings = append(
			findings,
			"Suspicious country indicator: Russia",
		)
	}

	// Brand impersonation detection
	for _, brand := range SuspiciousBrands {

		if strings.Contains(
			host,
			brand,
		) {

			findings = append(
				findings,
				"Brand impersonation detected: "+brand,
			)
		}
	}
	fmt.Println(
	"DNS Findings:",
	dnsFindings,
)

	return findings
}

func DetectLookalikeDomain(
	host string,
) []string {

	findings := []string{}

	host = strings.ToLower(host)

	lookalikes := map[string]string{
		"micr0soft": "microsoft",
		"g00gle":    "google",
		"arnazon":   "amazon",
		"paypa1":    "paypal",
		"app1e":     "apple",
	}

	for fake, brand := range lookalikes {

		if strings.Contains(
			host,
			fake,
		) {

			findings = append(
				findings,
				"Lookalike domain detected: "+brand,
			)
		}
	}

	return findings
}
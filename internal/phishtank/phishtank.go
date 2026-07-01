package phishtank

import (
	"bufio"
	"os"
	"strings"
)

var PhishTankDomains =
	LoadPhishTankFeed(
		"feeds/phishtank_feed.txt",
	)

func LoadPhishTankFeed(
	filePath string,
) []string {

	var domains []string

	file, err :=
		os.Open(
			filePath,
		)

	if err != nil {
		return domains
	}

	defer file.Close()

	scanner :=
		bufio.NewScanner(
			file,
		)

	for scanner.Scan() {

		domains = append(
			domains,
			strings.ToLower(
				scanner.Text(),
			),
		)
	}

	return domains
}

func CheckPhishTank(
	url string,
) []string {

	findings := []string{}

	url = strings.ToLower(
		url,
	)

	for _, domain := range PhishTankDomains {

		if strings.Contains(
			url,
			domain,
		) {

			findings = append(
				findings,
				"PhishTank hit: "+domain,
			)
		}
	}

	return findings
}
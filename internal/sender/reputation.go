package sender

import "strings"

func CheckSenderReputation(
	sender string,
) []string {

	findings := []string{}

	sender = strings.ToLower(sender)

	suspiciousSenders := []string{
		"evil",
		"phishing",
		"attacker",
		"hack",
		"fraud",
	}

	for _, keyword := range suspiciousSenders {

		if strings.Contains(
			sender,
			keyword,
		) {

			findings = append(
				findings,
				"Suspicious sender reputation detected",
			)

			break
		}
	}

	return findings
}
package sigma

import (
	"fmt"
	"os"
)

func GenerateRule(
	sender string,
	riskScore int,
	mitre string,
	filename string,
) error {

	rule := fmt.Sprintf(`
title: Phishing Email Detected

logsource:
  product: email

detection:
  selection:
    sender: %s
    risk_score: %d

condition: selection

level: high

tags:
  - %s
`,
		sender,
		riskScore,
		mitre,
	)

	return os.WriteFile(
		filename,
		[]byte(rule),
		0644,
	)
}
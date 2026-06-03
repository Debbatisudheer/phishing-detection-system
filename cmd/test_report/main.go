package main

import (
	"fmt"

	"phishing-platform/internal/report"
)

func main() {

	err := report.GenerateReport(
		"attacker@evil.com",
		"Microsoft Security Alert",
		[]string{
			"https://evil.com/login",
		},
		[]string{
			"SPF FAIL",
			"DKIM FAIL",
			"DMARC FAIL",
			"Known malicious domain detected",
		},
		775,
		"CRITICAL",
		"QUARANTINE",
		"T1566 - Phishing",
	)

	if err != nil {
		panic(err)
	}

	fmt.Println(
		"Investigation Report Generated",
	)
}
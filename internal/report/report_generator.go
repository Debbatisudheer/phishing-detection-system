package report

import (
	"fmt"
	"os"
)

func GenerateReport(
	sender string,
	subject string,
	urls []string,
	findings []string,
	riskScore int,
	riskLevel string,
	decision string,
	mitre string,
) error {

	report := fmt.Sprintf(
		`====================================
PHISHING INVESTIGATION REPORT
====================================

Sender:
%s

Subject:
%s

URLs:
%v

Findings:
%v

Risk Score:
%d

Risk Level:
%s

Decision:
%s

MITRE:
%s

====================================
`,
		sender,
		subject,
		urls,
		findings,
		riskScore,
		riskLevel,
		decision,
		mitre,
	)

	return os.WriteFile(
		"investigation_report.txt",
		[]byte(report),
		0644,
	)
}
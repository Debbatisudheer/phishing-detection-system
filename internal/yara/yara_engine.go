package yara

import (
	"strings"
)

func ScanContent(
	content string,
) []string {

	findings := []string{}

	content = strings.ToLower(content)

	rules := map[string]string{
		"powershell":   "YARA rule matched: PowerShell",
		"wscript":      "YARA rule matched: WScript",
		"createobject": "YARA rule matched: CreateObject",
		"autoopen":     "YARA rule matched: AutoOpen",
		"cmd.exe":      "YARA rule matched: CMD Execution",
		"http://":      "YARA rule matched: URL Indicator",
		"https://":     "YARA rule matched: URL Indicator",
	}

	for keyword, finding := range rules {

		if strings.Contains(
			content,
			keyword,
		) {

			findings = append(
				findings,
				finding,
			)
		}
	}

	return findings
}
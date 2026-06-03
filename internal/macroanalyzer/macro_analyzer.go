package macroanalyzer

import (
	"strings"
)

func AnalyzeMacroContent(
	content string,
) []string {

	var findings []string

	content = strings.ToLower(content)

	suspiciousMacros := []string{
		"autoopen",
		"document_open",
		"shell",
		"powershell",
		"createobject",
		"wscript",
		"cmd.exe",
	}

	for _, indicator := range suspiciousMacros {

		if strings.Contains(
			content,
			indicator,
		) {

			findings = append(
				findings,
				"Suspicious macro detected: "+indicator,
			)
		}
	}

	return findings
}
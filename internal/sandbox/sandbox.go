package sandbox

import (
	"strings"
)

func AnalyzeBehavior(
	content string,
) []string {

	var findings []string

	content = strings.ToLower(content)

	behaviors := map[string]string{
		"powershell":     "Sandbox behavior: PowerShell execution",
		"cmd.exe":        "Sandbox behavior: Command execution",
		"wscript":        "Sandbox behavior: Script execution",
		"createobject":   "Sandbox behavior: COM object creation",
		"downloadstring": "Sandbox behavior: Remote payload download",
		"http://":        "Sandbox behavior: Network communication",
		"https://":       "Sandbox behavior: Network communication",
	}

	for keyword, finding := range behaviors {

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
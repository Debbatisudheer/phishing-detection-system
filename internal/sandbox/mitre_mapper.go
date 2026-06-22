package sandbox

import "strings"

func MapSandboxMITRE(
	findings []string,
) string {

	for _, finding :=
		range findings {

		if strings.Contains(
			finding,
			"PowerShell",
		) {

			return "T1059.001"
		}

		if strings.Contains(
			finding,
			"URL",
		) {

			return "T1566.002"
		}
	}

	return ""
}
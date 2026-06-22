package sandbox

import (
	"regexp"
	"strings"
)

func DetectDroppedFiles(
	content string,
) []string {

	var findings []string

	re :=
		regexp.MustCompile(
			`(?i)[a-zA-Z0-9_\-]+\.(exe|dll|ps1|bat|vbs)`,
		)

	matches :=
		re.FindAllString(
			content,
			-1,
		)

	seen :=
		make(map[string]bool)

	systemProcesses :=
		map[string]bool{
			"powershell.exe": true,
			"cmd.exe":        true,
			"wscript.exe":    true,
			"cscript.exe":    true,
		}

	for _, file := range matches {

		file =
			strings.ToLower(
				file,
			)

		if seen[file] {
			continue
		}

		if systemProcesses[file] {
			continue
		}

		seen[file] = true

		findings = append(
			findings,
			"Dropped File Detected: "+file,
		)
	}

	return findings
}
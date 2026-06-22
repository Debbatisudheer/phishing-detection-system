package sandbox

func SimulateProcessTree(
	content string,
) []string {

	var findings []string

	if Contains(
		content,
		"powershell",
	) {

		findings = append(
			findings,
			"Process Tree: powershell.exe",
		)
	}

	if Contains(
		content,
		"invoke-webrequest",
	) {

		findings = append(
			findings,
			"Process Tree: powershell.exe -> network download",
		)
	}

	if Contains(
		content,
		".exe",
	) {

		findings = append(
			findings,
			"Process Tree: downloaded executable",
		)
	}

	return findings
}
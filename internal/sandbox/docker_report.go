package sandbox

func BuildDockerReport(
	output string,
	err error,
) []string {

	var findings []string

	findings = append(
		findings,
		"Docker Event: Container Started",
	)

	findings = append(
		findings,
		"Docker Event: File Mounted",
	)

	if err != nil {

		findings = append(
			findings,
			"Docker Event: Execution Failed",
		)

	} else {

		findings = append(
			findings,
			"Docker Event: Execution Successful",
		)
	}

	findings = append(
		findings,
		"Docker Event: Container Destroyed",
	)

	return findings
}
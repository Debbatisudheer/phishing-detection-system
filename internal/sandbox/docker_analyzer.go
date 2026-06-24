package sandbox

import "strings"

func AnalyzeDockerOutput(
	output string,
) []string {

	var findings []string

	lower :=
		strings.ToLower(
			output,
		)

	// PowerShell Detection
	if strings.Contains(
		lower,
		"powershell",
	) {

		findings = append(
			findings,
			"Docker Analysis: PowerShell Indicator",
		)
	}

	// URL Detection
	if strings.Contains(
		lower,
		"http",
	) {

		findings = append(
			findings,
			"Docker Analysis: URL Detected",
		)
	}

	// Executable Detection
	if strings.Contains(
		lower,
		".exe",
	) {

		findings = append(
			findings,
			"Docker Analysis: Executable Reference",
		)
	}

	// Hash Detection
	if strings.Contains(
		lower,
		"sha256",
	) {

		findings = append(
			findings,
			"Docker Analysis: Hash Generated",
		)
	}

	// Text File Detection
	if strings.Contains(
		lower,
		"ascii text",
	) ||
		strings.Contains(
			lower,
			"unicode text",
		) ||
		strings.Contains(
			lower,
			"utf-8",
		) {

		findings = append(
			findings,
			"Docker Analysis: Text File Detected",
		)
	}

	// ZIP Detection
	if strings.Contains(
		lower,
		"zip archive",
	) {

		findings = append(
			findings,
			"Docker Analysis: ZIP Archive Detected",
		)
	}

	// PDF Detection
	if strings.Contains(
		lower,
		"pdf document",
	) {

		findings = append(
			findings,
			"Docker Analysis: PDF Document Detected",
		)
	}

	// Base64 Detection
	if strings.Contains(
		lower,
		"frombase64string",
	) {

		findings = append(
			findings,
			"Docker Analysis: Base64 Decode Function",
		)
	}

	// Encoded PowerShell
	if strings.Contains(
		lower,
		"-enc",
	) ||
		strings.Contains(
			lower,
			"-encodedcommand",
		) {

		findings = append(
			findings,
			"Docker Analysis: Encoded PowerShell",
		)
	}

	// Download Activity
	if strings.Contains(
		lower,
		"invoke-webrequest",
	) {

		findings = append(
			findings,
			"Docker Analysis: Download Activity",
		)
	}

	// Registry Persistence
	if strings.Contains(
		lower,
		"currentversion\\run",
	) {

		findings = append(
			findings,
			"Docker Analysis: Registry Persistence",
		)
	}

	// ClamAV Results
	if strings.Contains(
		lower,
		"no supported database files found",
	) {

		findings = append(
			findings,
			"Docker Analysis: ClamAV Database Missing",
		)

	} else if strings.Contains(
		lower,
		"infected files: 0",
	) {

		findings = append(
			findings,
			"Docker Analysis: ClamAV Clean",
		)

	} else if strings.Contains(
		lower,
		"infected files: 1",
	) {

		findings = append(
			findings,
			"Docker Analysis: ClamAV Malware Detected",
		)
	}

	// ClamAV Timeout
	if strings.Contains(
		lower,
		"timed out",
	) ||
		strings.Contains(
			lower,
			"context deadline exceeded",
		) {

		findings = append(
			findings,
			"Docker Analysis: ClamAV Timeout",
		)
	}

	// YARA Matches
	if strings.Contains(
		lower,
		"powershell_downloader",
	) {

		findings = append(
			findings,
			"Docker YARA Match: PowerShell Downloader",
		)
	}

	if strings.Contains(
		lower,
		"registry_persistence",
	) {

		findings = append(
			findings,
			"Docker YARA Match: Registry Persistence",
		)
	}

	if strings.Contains(
		lower,
		"encoded_powershell",
	) {

		findings = append(
			findings,
			"Docker YARA Match: Encoded PowerShell",
		)
	}

	return findings
}
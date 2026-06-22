package sandbox

import (
	"regexp"
	"strings"
)

func ExtractSandboxURLs(
	content string,
) []string {

	regex :=
		regexp.MustCompile(
			`https?://[^\s"']+`,
		)

	return regex.FindAllString(
		content,
		-1,
	)
}

func ExtractSandboxIPs(
	content string,
) []string {

	regex :=
		regexp.MustCompile(
			`\b(?:\d{1,3}\.){3}\d{1,3}\b`,
		)

	return regex.FindAllString(
		content,
		-1,
	)
}

func ExtractSandboxEmails(
	content string,
) []string {

	regex :=
		regexp.MustCompile(
			`[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}`,
		)

	return regex.FindAllString(
		content,
		-1,
	)
}

func ExtractSandboxDomains(
	content string,
) []string {

	regex :=
		regexp.MustCompile(
			`\b[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}\b`,
		)

	matches :=
		regex.FindAllString(
			content,
			-1,
		)

	var domains []string

	blocked :=
		map[string]bool{
			"exe":  true,
			"dll":  true,
			"ps1":  true,
			"bat":  true,
			"vbs":  true,
			"docm": true,
			"xlsm": true,
		}

	seen :=
		make(map[string]bool)

	for _, domain := range matches {

		parts :=
			strings.Split(
				domain,
				".",
			)

		if len(parts) < 2 {
			continue
		}

		ext :=
			strings.ToLower(
				parts[len(parts)-1],
			)

		if blocked[ext] {
			continue
		}

		if seen[domain] {
			continue
		}

		seen[domain] = true

		domains = append(
			domains,
			domain,
		)
	}

	return domains
}
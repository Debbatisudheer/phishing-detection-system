package zipanalyzer

import (
	"archive/zip"
	"strings"
)

func DetectNestedZIP(
	zipPath string,
) []string {

	var findings []string

	reader, err :=
		zip.OpenReader(
			zipPath,
		)

	if err != nil {
		return findings
	}

	defer reader.Close()

	for _, file := range reader.File {

		if strings.HasSuffix(
			strings.ToLower(
				file.Name,
			),
			".zip",
		) {

			findings = append(
				findings,
				"Nested ZIP detected: "+file.Name,
			)
		}
	}

	return findings
}
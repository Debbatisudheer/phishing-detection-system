package macroanalyzer

import (
	"archive/zip"
	"io"
	"strings"
)

func ExtractMacroText(
	filePath string,
) string {

	zipReader, err :=
		zip.OpenReader(
			filePath,
		)

	if err != nil {
		return ""
	}

	defer zipReader.Close()

	for _, file := range zipReader.File {

		if strings.Contains(
			strings.ToLower(file.Name),
			"jdedata.bin",
		) {

			rc, err :=
				file.Open()

			if err != nil {
				return ""
			}

			defer rc.Close()

			data, err :=
				io.ReadAll(
					rc,
				)

			if err != nil {
				return ""
			}

			return strings.ToLower(
				string(data),
			)
		}
	}

	return ""
}
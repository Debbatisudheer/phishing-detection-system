package macroanalyzer

import (
	"archive/zip"
	"io"
	"strings"
)

func ExtractWPSMacroText(
	filePath string,
) string {

	reader, err :=
		zip.OpenReader(
			filePath,
		)

	if err != nil {
		return ""
	}

	defer reader.Close()

	for _, file := range reader.File {

		if file.Name ==
			"word/JDEData.bin" {

			rc, err :=
				file.Open()

			if err != nil {
				return ""
			}

			data, err :=
				io.ReadAll(
					rc,
				)

			rc.Close()

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
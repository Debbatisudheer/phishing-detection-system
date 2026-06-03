package zipanalyzer

import (
	"archive/zip"
)

func ExtractZIPContents(
	zipPath string,
) ([]string, error) {

	var files []string

	reader, err :=
		zip.OpenReader(
			zipPath,
		)

	if err != nil {
		return nil, err
	}

	defer reader.Close()

	for _, file := range reader.File {

		files = append(
			files,
			file.Name,
		)
	}

	return files, nil
}
package zipanalyzer

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
)

func ExtractArtifactsForSandbox(
	zipPath string,
) ([]string, error) {

	var extractedFiles []string

	err := os.MkdirAll(
		"sandbox_files",
		0755,
	)

	if err != nil {
		return nil, err
	}

	reader, err :=
		zip.OpenReader(
			zipPath,
		)

	if err != nil {
		return nil, err
	}

	defer reader.Close()

	for _, file := range reader.File {

		source, err :=
			file.Open()

		if err != nil {
			continue
		}

		targetPath :=
			filepath.Join(
				"sandbox_files",
				file.Name,
			)

		target, err :=
			os.Create(
				targetPath,
			)

		if err != nil {

			source.Close()
			continue
		}

		io.Copy(
			target,
			source,
		)

		target.Close()
		source.Close()

		extractedFiles = append(
			extractedFiles,
			targetPath,
		)
	}

	return extractedFiles, nil
}
package zipanalyzer

import (
	"archive/zip"
	"os"
	"testing"
)

func TestExtractArtifactsForSandbox(t *testing.T) {

	tmp := "artifacts.zip"

	file, err := os.Create(tmp)
	if err != nil {
		t.Fatal(err)
	}
	file.Close()

	zipFile, err := os.Create(tmp)
	if err != nil {
		t.Fatal(err)
	}

	writer := zip.NewWriter(zipFile)

	f, _ := writer.Create("invoice.docm")
	f.Write([]byte("macro"))

	f, _ = writer.Create("payload.exe")
	f.Write([]byte("exe"))

	writer.Close()
	zipFile.Close()

	defer os.Remove(tmp)
	defer os.RemoveAll("sandbox_files")

	files, err := ExtractArtifactsForSandbox(tmp)

	if err != nil {
		t.Fatal(err)
	}

	if len(files) != 2 {
		t.Fatalf(
			"expected 2 extracted files got %d",
			len(files),
		)
	}

	for _, file := range files {

		if _, err := os.Stat(file); err != nil {
			t.Errorf(
				"expected extracted file %s",
				file,
			)
		}
	}
}
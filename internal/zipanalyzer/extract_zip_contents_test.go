package zipanalyzer

import (
	"archive/zip"
	"os"
	"reflect"
	"testing"
)

func TestExtractZIPContents(t *testing.T) {

	tmp := "test.zip"

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

	result, err := ExtractZIPContents(tmp)

	if err != nil {
		t.Fatal(err)
	}

	expected := []string{
		"invoice.docm",
		"payload.exe",
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf(
			"expected %v got %v",
			expected,
			result,
		)
	}
}
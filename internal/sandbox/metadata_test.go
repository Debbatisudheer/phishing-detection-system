package sandbox

import (
	"os"
	"testing"
)

func TestGetFileMetadata(t *testing.T) {

	content := []byte("Hello Sandbox")

	tmp, err := os.CreateTemp("", "sandbox-*.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmp.Name())

	if _, err := tmp.Write(content); err != nil {
		t.Fatal(err)
	}

	tmp.Close()

	meta := GetFileMetadata(tmp.Name())

	if meta.FileName == "" {
		t.Error("expected filename")
	}

	if meta.FileSize == 0 {
		t.Error("expected filesize")
	}

	if meta.Extension != ".txt" {
		t.Errorf("expected .txt got %s", meta.Extension)
	}

	if meta.MIMEType == "" {
		t.Error("expected mimetype")
	}

	if meta.MD5 == "" {
		t.Error("expected md5")
	}
}
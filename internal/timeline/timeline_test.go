package timeline

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func TestLogEvent(t *testing.T) {

	// Backup stdout
	old := os.Stdout

	r, w, _ := os.Pipe()
	os.Stdout = w

	LogEvent("Email Parsed")

	w.Close()

	os.Stdout = old

	var buf bytes.Buffer
	io.Copy(&buf, r)

	output := buf.String()

	if !strings.Contains(output, "Email Parsed") {
		t.Fatalf(
			"expected output to contain event, got %q",
			output,
		)
	}

	if !strings.Contains(output, "[") ||
		!strings.Contains(output, "]") {

		t.Fatal(
			"timestamp format missing",
		)
	}
}
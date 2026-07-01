package pdfanalyzer

import "testing"

func TestExtractPDFText_InvalidFile(t *testing.T) {

	result := ExtractPDFText("does_not_exist.pdf")

	if result != "" {
		t.Errorf("expected empty string got %q", result)
	}
}
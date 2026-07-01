package macroanalyzer

import "testing"

func TestExtractMacroText_InvalidFile(t *testing.T) {

	result := ExtractMacroText("does_not_exist.docm")

	if result != "" {
		t.Errorf("expected empty string got %q", result)
	}
}

func TestExtractWPSMacroText_InvalidFile(t *testing.T) {

	result := ExtractWPSMacroText("does_not_exist.docm")

	if result != "" {
		t.Errorf("expected empty string got %q", result)
	}
}
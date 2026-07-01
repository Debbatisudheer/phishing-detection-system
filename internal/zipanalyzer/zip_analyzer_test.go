package zipanalyzer

import (
	"reflect"
	"testing"
)

func TestAnalyzeZIPContents(t *testing.T) {

	tests := []struct {
		name     string
		file     string
		expected []string
	}{
		{
			name: "ZIP File",
			file: "invoice.zip",
			expected: []string{
				"ZIP contains suspicious file: invoice.exe",
				"ZIP contains suspicious file: payload.ps1",
			},
		},
		{
			name:     "PDF",
			file:     "manual.pdf",
			expected: []string{},
		},
	}

	for _, tc := range tests {

		t.Run(tc.name, func(t *testing.T) {

			result := AnalyzeZIPContents(tc.file)

			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("expected %v got %v", tc.expected, result)
			}
		})
	}
}
package zipanalyzer

import (
	"reflect"
	"testing"
)

func TestAnalyzeZIPFileContents(t *testing.T) {

	tests := []struct {
		name     string
		files    []string
		expected []string
	}{
		{
			name: "Executable",
			files: []string{
				"virus.exe",
			},
			expected: []string{
				"ZIP contains executable file: virus.exe",
			},
		},
		{
			name: "PowerShell",
			files: []string{
				"payload.ps1",
			},
			expected: []string{
				"ZIP contains PowerShell file: payload.ps1",
			},
		},
		{
			name: "Macro Word",
			files: []string{
				"invoice.docm",
			},
			expected: []string{
				"ZIP contains macro-enabled document: invoice.docm",
			},
		},
		{
			name: "Macro Excel",
			files: []string{
				"report.xlsm",
			},
			expected: []string{
				"ZIP contains macro-enabled spreadsheet: report.xlsm",
			},
		},
		{
			name: "Safe",
			files: []string{
				"notes.txt",
			},
			expected: []string{},
		},
	}

	for _, tc := range tests {

		t.Run(tc.name, func(t *testing.T) {

			result := AnalyzeZIPFileContents(tc.files)

			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("expected %v got %v", tc.expected, result)
			}
		})
	}
}
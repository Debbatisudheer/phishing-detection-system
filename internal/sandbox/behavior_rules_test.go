package sandbox

import (
	"reflect"
	"testing"
)

func TestAnalyzeBehaviorRules(t *testing.T) {

	tests := []struct {
		name     string
		findings []string
		expected []string
	}{
		{
			name: "PowerShell Downloader",
			findings: []string{
				"YARA rule matched: PowerShell",
				"YARA rule matched: URL Indicator",
			},
			expected: []string{
				"Behavior Rule: PowerShell Downloader Detected",
			},
		},
		{
			name: "Malicious Office",
			findings: []string{
				"Macro Document Detected",
				"YARA rule matched: PowerShell",
			},
			expected: []string{
				"Behavior Rule: Malicious Office Execution",
			},
		},
		{
			name: "Malware Dropper",
			findings: []string{
				"Executable Detected",
				"YARA rule matched: PowerShell",
			},
			expected: []string{
				"Behavior Rule: Malware Dropper Detected",
			},
		},
		{
			name: "Multiple Rules",
			findings: []string{
				"Executable Detected",
				"Macro Document Detected",
				"YARA rule matched: PowerShell",
				"YARA rule matched: URL Indicator",
			},
			expected: []string{
				"Behavior Rule: PowerShell Downloader Detected",
				"Behavior Rule: Malicious Office Execution",
				"Behavior Rule: Malware Dropper Detected",
			},
		},
		{
			name: "No Match",
			findings: []string{
				"Nothing",
			},
			expected: nil,
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {

			result := AnalyzeBehaviorRules(tt.findings)

			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf(
					"expected %v got %v",
					tt.expected,
					result,
				)
			}
		})
	}
}
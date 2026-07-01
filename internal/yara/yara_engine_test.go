package yara

import (
	"reflect"
	"sort"
	"testing"
)

func TestScanContent(t *testing.T) {

	tests := []struct {
		name     string
		content  string
		expected []string
	}{
		{
			name:    "PowerShell",
			content: "powershell -enc abc",
			expected: []string{
				"YARA rule matched: PowerShell",
			},
		},
		{
			name:    "WScript",
			content: "WScript.Shell",
			expected: []string{
				"YARA rule matched: WScript",
			},
		},
		{
			name:    "CreateObject",
			content: "CreateObject(\"WScript.Shell\")",
			expected: []string{
				"YARA rule matched: WScript",
				"YARA rule matched: CreateObject",
			},
		},
		{
			name:    "AutoOpen",
			content: "AutoOpen",
			expected: []string{
				"YARA rule matched: AutoOpen",
			},
		},
		{
			name:    "CMD",
			content: "cmd.exe /c calc.exe",
			expected: []string{
				"YARA rule matched: CMD Execution",
			},
		},
		{
			name:    "HTTP URL",
			content: "http://evil.com",
			expected: []string{
				"YARA rule matched: URL Indicator",
			},
		},
		{
			name:    "HTTPS URL",
			content: "https://evil.com",
			expected: []string{
				"YARA rule matched: URL Indicator",
			},
		},
		{
			name:    "Multiple Rules",
			content: "powershell AutoOpen https://evil.com",
			expected: []string{
				"YARA rule matched: PowerShell",
				"YARA rule matched: AutoOpen",
				"YARA rule matched: URL Indicator",
			},
		},
		{
			name:     "Safe Content",
			content:  "Hello world",
			expected: []string{},
		},
		{
			name:     "Empty",
			content:  "",
			expected: []string{},
		},
	}

	for _, tc := range tests {

		t.Run(tc.name, func(t *testing.T) {

			result := ScanContent(tc.content)

			sort.Strings(result)
sort.Strings(tc.expected)

if !reflect.DeepEqual(result, tc.expected) {
	t.Errorf(
		"expected %v got %v",
		tc.expected,
		result,
	)
}
		})
	}
}
package sandbox

import (
	"reflect"
	"sort"
	"testing"
)

func TestAnalyzeBehavior(t *testing.T) {

	tests := []struct {
		name     string
		content  string
		expected []string
	}{
		{
			name:    "PowerShell",
			content: "PowerShell",
			expected: []string{
				"Sandbox behavior: PowerShell execution",
			},
		},
		{
			name:    "CMD",
			content: "cmd.exe",
			expected: []string{
				"Sandbox behavior: Command execution",
			},
		},
		{
			name:    "WScript",
			content: "wscript",
			expected: []string{
				"Sandbox behavior: Script execution",
			},
		},
		{
			name:    "CreateObject",
			content: "CreateObject",
			expected: []string{
				"Sandbox behavior: COM object creation",
			},
		},
		{
			name:    "DownloadString",
			content: "DownloadString",
			expected: []string{
				"Sandbox behavior: Remote payload download",
			},
		},
		{
			name:    "HTTP",
			content: "http://evil.com",
			expected: []string{
				"Sandbox behavior: Network communication",
			},
		},
		{
			name:    "HTTPS",
			content: "https://evil.com",
			expected: []string{
				"Sandbox behavior: Network communication",
			},
		},
		{
			name: "Multiple",
			content: "PowerShell cmd.exe wscript CreateObject DownloadString https://evil.com",
			expected: []string{
				"Sandbox behavior: PowerShell execution",
				"Sandbox behavior: Command execution",
				"Sandbox behavior: Script execution",
				"Sandbox behavior: COM object creation",
				"Sandbox behavior: Remote payload download",
				"Sandbox behavior: Network communication",
			},
		},
		{
			name:     "Safe",
			content:  "Hello World",
			expected: nil,
		},
		{
			name:     "Empty",
			content:  "",
			expected: nil,
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {

			result := AnalyzeBehavior(tt.content)

			sort.Strings(result)
sort.Strings(tt.expected)

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
package macroanalyzer

import (
	"reflect"
	"testing"
)

func TestAnalyzeMacroContent(t *testing.T) {

	tests := []struct {
		name     string
		content  string
		expected []string
	}{
		{
			name:    "AutoOpen",
			content: "AutoOpen",
			expected: []string{
				"Suspicious macro detected: autoopen",
			},
		},
		{
			name:    "PowerShell",
			content: "powershell -enc abc",
			expected: []string{
    "Suspicious macro detected: shell",
    "Suspicious macro detected: powershell",
},
		},
		{
			name:    "Shell",
			content: "Shell(\"cmd.exe\")",
			expected: []string{
				"Suspicious macro detected: shell",
				"Suspicious macro detected: cmd.exe",
			},
		},
		{
			name:    "CreateObject",
			content: "CreateObject(\"WScript.Shell\")",
			expected: []string{
    "Suspicious macro detected: shell",
    "Suspicious macro detected: createobject",
    "Suspicious macro detected: wscript",
},
		},
		{
			name:    "Multiple Indicators",
			content: "AutoOpen powershell cmd.exe",
			expected: []string{
    "Suspicious macro detected: autoopen",
    "Suspicious macro detected: shell",
    "Suspicious macro detected: powershell",
    "Suspicious macro detected: cmd.exe",
},
		},
		{
			name:     "Safe Macro",
			content:  "Hello World",
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

			result := AnalyzeMacroContent(tc.content)

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
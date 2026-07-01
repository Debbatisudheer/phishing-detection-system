package sandbox

import (
	"reflect"
	"testing"
)

func TestSimulateProcessTree(t *testing.T) {

	tests := []struct {
		name     string
		content  string
		expected []string
	}{
		{
			name:    "PowerShell",
			content: "powershell",
			expected: []string{
				"Process Tree: powershell.exe",
			},
		},
		{
			name:    "Download",
			content: "powershell Invoke-WebRequest",
			expected: []string{
				"Process Tree: powershell.exe",
				"Process Tree: powershell.exe -> network download",
			},
		},
		{
			name:    "Executable",
			content: "payload.exe",
			expected: []string{
				"Process Tree: downloaded executable",
			},
		},
		{
			name:    "Everything",
			content: "powershell Invoke-WebRequest payload.exe",
			expected: []string{
				"Process Tree: powershell.exe",
				"Process Tree: powershell.exe -> network download",
				"Process Tree: downloaded executable",
			},
		},
		{
			name:     "No Match",
			content:  "hello world",
			expected: nil,
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {

			result := SimulateProcessTree(tt.content)

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
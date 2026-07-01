package sandbox

import (
	"reflect"
	"testing"
)

func TestDetectDroppedFiles(t *testing.T) {

	tests := []struct {
		name     string
		content  string
		expected []string
	}{
		{
			name:    "Executable",
			content: "invoice.exe",
			expected: []string{
				"Dropped File Detected: invoice.exe",
			},
		},
		{
			name:    "DLL",
			content: "payload.dll",
			expected: []string{
				"Dropped File Detected: payload.dll",
			},
		},
		{
			name:    "PowerShell Script",
			content: "payload.ps1",
			expected: []string{
				"Dropped File Detected: payload.ps1",
			},
		},
		{
			name:    "Batch File",
			content: "run.bat",
			expected: []string{
				"Dropped File Detected: run.bat",
			},
		},
		{
			name:    "VBS Script",
			content: "script.vbs",
			expected: []string{
				"Dropped File Detected: script.vbs",
			},
		},
		{
			name: "Multiple Files",
			content: `
invoice.exe
payload.dll
payload.ps1
run.bat
script.vbs
`,
			expected: []string{
				"Dropped File Detected: invoice.exe",
				"Dropped File Detected: payload.dll",
				"Dropped File Detected: payload.ps1",
				"Dropped File Detected: run.bat",
				"Dropped File Detected: script.vbs",
			},
		},
		{
			name:    "System Process Ignored",
			content: "powershell.exe cmd.exe wscript.exe cscript.exe",
			expected: nil,
		},
		{
			name: "Duplicate Files",
			content: `
invoice.exe
invoice.exe
invoice.exe
`,
			expected: []string{
				"Dropped File Detected: invoice.exe",
			},
		},
		{
			name:     "Safe Content",
			content:  "hello world",
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

			result := DetectDroppedFiles(tt.content)

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
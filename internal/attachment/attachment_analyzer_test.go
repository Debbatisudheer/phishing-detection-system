package attachment

import (
	"reflect"
	"testing"
)

func TestAnalyzeAttachments(t *testing.T) {

	tests := []struct {
		name        string
		attachments []string
		expected    []string
	}{
		{
			name: "ZIP File",
			attachments: []string{
				"invoice.zip",
			},
			expected: []string{
				"ZIP attachment detected",
			},
		},
		{
			name: "RAR File",
			attachments: []string{
				"archive.rar",
			},
			expected: []string{
				"RAR attachment detected",
			},
		},
		{
			name: "DOCM File",
			attachments: []string{
				"invoice.docm",
			},
			expected: []string{
				"Macro-enabled Office document detected",
			},
		},
		{
			name: "XLSM File",
			attachments: []string{
				"report.xlsm",
			},
			expected: []string{
				"Macro-enabled Excel document detected",
			},
		},
		{
			name: "PDF File",
			attachments: []string{
				"manual.pdf",
			},
			expected: []string{
				"PDF attachment detected",
			},
		},
		{
			name: "EXE File",
			attachments: []string{
				"malware.exe",
			},
			expected: []string{
				"Suspicious attachment detected: malware.exe",
			},
		},
		{
			name: "PowerShell File",
			attachments: []string{
				"payload.ps1",
			},
			expected: []string{
				"Suspicious attachment detected: payload.ps1",
			},
		},
		{
			name: "JavaScript File",
			attachments: []string{
				"attack.js",
			},
			expected: []string{
				"Suspicious attachment detected: attack.js",
			},
		},
		{
			name: "Multiple Attachments",
			attachments: []string{
				"invoice.zip",
				"macro.docm",
				"virus.exe",
			},
			expected: []string{
				"ZIP attachment detected",
				"Macro-enabled Office document detected",
				"Suspicious attachment detected: virus.exe",
			},
		},
		{
			name: "Uppercase Extensions",
			attachments: []string{
				"FILE.ZIP",
				"DOCUMENT.PDF",
				"VIRUS.EXE",
			},
			expected: []string{
				"ZIP attachment detected",
				"PDF attachment detected",
				"Suspicious attachment detected: virus.exe",
			},
		},
		{
			name: "Safe File",
			attachments: []string{
				"notes.txt",
			},
			expected: []string{},
		},
		{
			name:        "No Attachments",
			attachments: []string{},
			expected:    []string{},
		},
	}

	for _, tc := range tests {

		t.Run(tc.name, func(t *testing.T) {

			result := AnalyzeAttachments(tc.attachments)

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
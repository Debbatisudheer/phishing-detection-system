package sandbox

import (
	"strings"
	"testing"
)

func TestAnalyzeDockerOutput(t *testing.T) {

	tests := []struct {
		name     string
		output   string
		expected []string
	}{
		{
			name: "PowerShell",
			output: "powershell",
			expected: []string{
				"Docker Analysis: PowerShell Indicator",
			},
		},
		{
			name: "URL",
			output: "https://evil.com",
			expected: []string{
				"Docker Analysis: URL Detected",
			},
		},
		{
			name: "Executable",
			output: "payload.exe",
			expected: []string{
				"Docker Analysis: Executable Reference",
			},
		},
		{
			name: "Hash",
			output: "sha256",
			expected: []string{
				"Docker Analysis: Hash Generated",
			},
		},
		{
			name: "Text",
			output: "ASCII text",
			expected: []string{
				"Docker Analysis: Text File Detected",
			},
		},
		{
			name: "ZIP",
			output: "Zip archive",
			expected: []string{
				"Docker Analysis: ZIP Archive Detected",
			},
		},
		{
			name: "PDF",
			output: "PDF document",
			expected: []string{
				"Docker Analysis: PDF Document Detected",
			},
		},
		{
			name: "Base64",
			output: "FromBase64String",
			expected: []string{
				"Docker Analysis: Base64 Decode Function",
			},
		},
		{
			name: "Encoded PS",
			output: "-EncodedCommand",
			expected: []string{
				"Docker Analysis: Encoded PowerShell",
			},
		},
		{
			name: "Download",
			output: "Invoke-WebRequest",
			expected: []string{
				"Docker Analysis: Download Activity",
			},
		},
		{
			name: "Registry",
			output: "CurrentVersion\\Run",
			expected: []string{
				"Docker Analysis: Registry Persistence",
			},
		},
		{
			name: "ClamAV Clean",
			output: "Infected files: 0",
			expected: []string{
				"Docker Analysis: ClamAV Clean",
			},
		},
		{
			name: "ClamAV Malware",
			output: "Infected files: 1",
			expected: []string{
				"Docker Analysis: ClamAV Malware Detected",
			},
		},
		{
			name: "Timeout",
			output: "context deadline exceeded",
			expected: []string{
				"Docker Analysis: ClamAV Timeout",
			},
		},
		{
    name: "YARA PowerShell",
    output: "powershell_downloader",
    expected: []string{
        "Docker Analysis: PowerShell Indicator",
        "Docker YARA Match: PowerShell Downloader",
    },
},
		{
			name: "YARA Registry",
			output: "registry_persistence",
			expected: []string{
				"Docker YARA Match: Registry Persistence",
			},
		},
		{
    name: "YARA Encoded",
    output: "encoded_powershell",
    expected: []string{
        "Docker Analysis: PowerShell Indicator",
        "Docker YARA Match: Encoded PowerShell",
    },
},
		{
			name:     "Safe",
			output:   "hello world",
			expected: nil,
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {

			result := AnalyzeDockerOutput(tt.output)

			if len(result) != len(tt.expected) {
				t.Fatalf("expected %d findings got %d",
					len(tt.expected),
					len(result))
			}

			for i := range result {
				if !strings.Contains(result[i], tt.expected[i]) {
					t.Errorf(
						"expected %q got %q",
						tt.expected[i],
						result[i],
					)
				}
			}
		})
	}
}
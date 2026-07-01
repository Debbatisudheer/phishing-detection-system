package sandbox

import "testing"

func TestMapSandboxMITRE(t *testing.T) {

	tests := []struct {
		name     string
		findings []string
		expected string
	}{
		{
			name: "PowerShell",
			findings: []string{
				"Sandbox behavior: PowerShell execution",
			},
			expected: "T1059.001",
		},
		{
			name: "URL",
			findings: []string{
				"Sandbox IOC URL: https://evil.com",
			},
			expected: "T1566.002",
		},
		{
			name: "PowerShell Takes Priority",
			findings: []string{
				"Sandbox behavior: PowerShell execution",
				"Sandbox IOC URL: https://evil.com",
			},
			expected: "T1059.001",
		},
		{
			name: "No Match",
			findings: []string{
				"Safe file",
			},
			expected: "",
		},
		{
			name:     "Empty",
			findings: []string{},
			expected: "",
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {

			result := MapSandboxMITRE(tt.findings)

			if result != tt.expected {

				t.Errorf(
					"expected %s got %s",
					tt.expected,
					result,
				)
			}
		})
	}
}
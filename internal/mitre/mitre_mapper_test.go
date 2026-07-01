package mitre

import (
	"reflect"
	"sort"
	"testing"
)

func TestMapTechnique(t *testing.T) {

	tests := []struct {
		name     string
		subject  string
		body     string
		expected string
	}{
		{
			name:     "HTTP Link",
			subject:  "Invoice",
			body:     "Click http://evil.com",
			expected: "T1566 - Phishing",
		},
		{
			name:     "HTTPS Link",
			subject:  "Invoice",
			body:     "Visit https://evil.com",
			expected: "T1566 - Phishing",
		},
		{
			name:     "Login",
			subject:  "",
			body:     "Please login now",
			expected: "T1204 - User Execution",
		},
		{
			name:     "Reset Password",
			subject:  "",
			body:     "Reset password immediately",
			expected: "T1204 - User Execution",
		},
		{
			name:     "No Match",
			subject:  "Meeting",
			body:     "See you tomorrow",
			expected: "NO_MITRE_MATCH",
		},
	}

	for _, tc := range tests {

		t.Run(tc.name, func(t *testing.T) {

			result := MapTechnique(
				tc.subject,
				tc.body,
			)

			if result != tc.expected {

				t.Errorf(
					"expected %s got %s",
					tc.expected,
					result,
				)
			}
		})
	}
}

func TestMapFileTechniques(t *testing.T) {

	tests := []struct {
		name     string
		findings []string
		expected []string
	}{
		{
			name: "PowerShell",
			findings: []string{
				"PowerShell detected",
			},
			expected: []string{
				"T1059.001 - PowerShell",
			},
		},
		{
			name: "Macro",
			findings: []string{
				"Macro-enabled Office document detected",
			},
			expected: []string{
				"T1566.001 - Spearphishing Attachment",
			},
		},
		{
			name: "URL",
			findings: []string{
				"Suspicious URL",
			},
			expected: []string{
				"T1566.002 - Spearphishing Link",
			},
		},
		{
			name: "Multiple",
			findings: []string{
				"PowerShell detected",
				"Macro detected",
				"URL detected",
			},
			expected: []string{
				"T1059.001 - PowerShell",
				"T1566.001 - Spearphishing Attachment",
				"T1566.002 - Spearphishing Link",
			},
		},
		{
			name: "Duplicate",
			findings: []string{
				"PowerShell",
				"PowerShell",
			},
			expected: []string{
				"T1059.001 - PowerShell",
			},
		},
		{
			name: "No Match",
			findings: []string{
				"Safe document",
			},
			expected: []string{},
		},
	}

	for _, tc := range tests {

		t.Run(tc.name, func(t *testing.T) {

			result := MapFileTechniques(tc.findings)

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
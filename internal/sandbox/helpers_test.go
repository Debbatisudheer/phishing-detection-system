package sandbox

import "testing"

func TestContains(t *testing.T) {

	tests := []struct {
		name     string
		content  string
		value    string
		expected bool
	}{
		{
			name:     "Exact Match",
			content:  "PowerShell",
			value:    "PowerShell",
			expected: true,
		},
		{
			name:     "Case Insensitive",
			content:  "PoWeRsHeLl",
			value:    "powershell",
			expected: true,
		},
		{
			name:     "Substring",
			content:  "Invoke PowerShell Script",
			value:    "powershell",
			expected: true,
		},
		{
			name:     "No Match",
			content:  "Hello World",
			value:    "powershell",
			expected: false,
		},
		{
			name:     "Empty Content",
			content:  "",
			value:    "powershell",
			expected: false,
		},
		{
			name:     "Empty Search",
			content:  "PowerShell",
			value:    "",
			expected: true,
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {

			result := Contains(
				tt.content,
				tt.value,
			)

			if result != tt.expected {
				t.Errorf(
					"expected %v got %v",
					tt.expected,
					result,
				)
			}
		})
	}
}
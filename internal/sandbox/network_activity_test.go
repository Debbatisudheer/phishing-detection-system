package sandbox

import (
	"reflect"
	"testing"
)

func TestAnalyzeNetworkActivity(t *testing.T) {

	tests := []struct {
		name     string
		content  string
		expected []string
	}{
		{
			name:    "HTTP",
			content: "http://evil.com",
			expected: []string{
				"Network Activity: Outbound Connection Detected",
				"Network Activity: Internet Communication",
			},
		},
		{
			name:    "HTTPS",
			content: "https://evil.com",
			expected: []string{
				"Network Activity: Outbound Connection Detected",
				"Network Activity: Internet Communication",
			},
		},
		{
			name:    "Invoke-WebRequest",
			content: "Invoke-WebRequest https://evil.com",
			expected: []string{
				"Network Activity: Outbound Connection Detected",
				"Network Activity: Internet Communication",
				"Network Activity: Potential Payload Download",
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

			result := AnalyzeNetworkActivity(tt.content)

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
package sandbox

import (
	"reflect"
	"testing"
)

func TestDetectPersistence(t *testing.T) {

	tests := []struct {
		name     string
		content  string
		expected []string
	}{
		{
			name:    "Registry Run Key",
			content: `CurrentVersion\Run`,
			expected: []string{
				"Persistence Detected: Registry Run Key Modification",
			},
		},
		{
			name:    "New Item Property",
			content: "New-ItemProperty",
			expected: []string{
				"Persistence Detected: Registry Persistence Creation",
			},
		},
		{
			name:    "Set Item Property",
			content: "Set-ItemProperty",
			expected: []string{
				"Persistence Detected: Registry Value Modification",
			},
		},
		{
			name: "Multiple",
			content: `CurrentVersion\Run
New-ItemProperty
Set-ItemProperty`,
			expected: []string{
				"Persistence Detected: Registry Run Key Modification",
				"Persistence Detected: Registry Persistence Creation",
				"Persistence Detected: Registry Value Modification",
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

			result := DetectPersistence(tt.content)

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
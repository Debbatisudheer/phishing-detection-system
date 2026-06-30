package parser

import (
	"reflect"
	"testing"
)

func TestExtractURLs(t *testing.T) {

	tests := []struct {
		name     string
		input    string
		expected []string
	}{
		{
			name:  "Single URL",
			input: "Visit https://google.com",
			expected: []string{
				"https://google.com",
			},
		},
		{
			name:  "Multiple URLs",
			input: "https://google.com https://github.com",
			expected: []string{
				"https://google.com",
				"https://github.com",
			},
		},
		{
			name:     "No URL",
			input:    "Hello World",
			expected: []string{},
		},
	}

	for _, tc := range tests {

		t.Run(tc.name, func(t *testing.T) {

			result := ExtractURLs(tc.input)

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
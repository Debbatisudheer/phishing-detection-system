package threatfeed

import (
	"reflect"
	"testing"
)

func TestCheckThreatFeed(t *testing.T) {

	// Override global feed for testing
	MaliciousDomains = []string{
		"evil.com",
		"phishing.com",
	}

	tests := []struct {
		name     string
		url      string
		expected []string
	}{
		{
			name: "Known Domain",
			url:  "https://evil.com/login",
			expected: []string{
				"Threat feed hit: evil.com",
			},
		},
		{
			name: "Second Domain",
			url:  "https://phishing.com/index.html",
			expected: []string{
				"Threat feed hit: phishing.com",
			},
		},
		{
			name:     "Safe Domain",
			url:      "https://google.com",
			expected: []string{},
		},
		{
			name:     "Empty URL",
			url:      "",
			expected: []string{},
		},
	}

	for _, tc := range tests {

		t.Run(tc.name, func(t *testing.T) {

			result := CheckThreatFeed(tc.url)

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
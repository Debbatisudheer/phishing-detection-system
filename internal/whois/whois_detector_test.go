package whois

import (
	"reflect"
	"testing"
)

func TestAnalyzeDomain(t *testing.T) {

	tests := []struct {
		name     string
		domain   string
		expected []string
	}{
		{
			name:   "XYZ Domain",
			domain: "evil.xyz",
			expected: []string{
				"WHOIS suspicious TLD detected: .xyz",
			},
		},
		{
			name:   "TOP Domain",
			domain: "attack.top",
			expected: []string{
				"WHOIS suspicious TLD detected: .top",
			},
		},
		{
			name:   "CLICK Domain",
			domain: "login.click",
			expected: []string{
				"WHOIS suspicious TLD detected: .click",
			},
		},
		{
			name:   "SHOP Domain",
			domain: "amazon.shop",
			expected: []string{
				"WHOIS suspicious TLD detected: .shop",
			},
		},
		{
			name:   "ONLINE Domain",
			domain: "secure.online",
			expected: []string{
				"WHOIS suspicious TLD detected: .online",
			},
		},
		{
			name:   "SITE Domain",
			domain: "fake.site",
			expected: []string{
				"WHOIS suspicious TLD detected: .site",
			},
		},
		{
			name:   "Case Insensitive",
			domain: "EVIL.XYZ",
			expected: []string{
				"WHOIS suspicious TLD detected: .xyz",
			},
		},
		{
			name:     "Safe Domain",
			domain:   "google.com",
			expected: []string{},
		},
		{
			name:     "Empty Domain",
			domain:   "",
			expected: []string{},
		},
	}

	for _, tc := range tests {

		t.Run(tc.name, func(t *testing.T) {

			result := AnalyzeDomain(tc.domain)

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
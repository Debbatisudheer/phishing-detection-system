package dnsreputation

import (
	"reflect"
	"testing"
)

func TestCheckDNSReputation(t *testing.T) {

	tests := []struct {
		name     string
		domain   string
		expected []string
	}{
		{
			name:   "Evil Domain",
			domain: "evil.com",
			expected: []string{
				"DNS reputation hit: evil.com",
			},
		},
		{
			name:   "Malware Domain",
			domain: "malware.com",
			expected: []string{
				"DNS reputation hit: malware.com",
			},
		},
		{
			name:   "Phishing Domain",
			domain: "phishing-site.xyz",
			expected: []string{
				"DNS reputation hit: phishing-site.xyz",
			},
		},
		{
			name:   "Subdomain Match",
			domain: "login.evil.com",
			expected: []string{
				"DNS reputation hit: evil.com",
			},
		},
		{
			name:   "Case Insensitive",
			domain: "EVIL.COM",
			expected: []string{
				"DNS reputation hit: evil.com",
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

			result := CheckDNSReputation(tc.domain)

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
package domain

import (
	"reflect"
	"testing"
)

func TestCheckDomainAge(t *testing.T) {

	tests := []struct {
		name     string
		host     string
		expected []string
	}{
		{
			name: "New Domain",
			host: "security-login.com",
			expected: []string{
				"Newly registered domain detected",
			},
		},
		{
			name:     "Normal Domain",
			host:     "google.com",
			expected: []string{},
		},
	}

	for _, tc := range tests {

		t.Run(tc.name, func(t *testing.T) {

			result := CheckDomainAge(tc.host)

			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("expected %v got %v", tc.expected, result)
			}
		})
	}
}

func TestDetectLookalikeDomain(t *testing.T) {

	tests := []struct {
		name     string
		host     string
		expected []string
	}{
		{
			name: "Fake Microsoft",
			host: "micr0soft-login.com",
			expected: []string{
				"Lookalike domain detected: microsoft",
			},
		},
		{
			name: "Fake Google",
			host: "g00gle.com",
			expected: []string{
				"Lookalike domain detected: google",
			},
		},
		{
			name:     "Legitimate",
			host:     "google.com",
			expected: []string{},
		},
	}

	for _, tc := range tests {

		t.Run(tc.name, func(t *testing.T) {

			result := DetectLookalikeDomain(tc.host)

			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("expected %v got %v", tc.expected, result)
			}
		})
	}
}

func TestDetectHomographDomain(t *testing.T) {

	tests := []struct {
		name     string
		host     string
		expected []string
	}{
		{
			name: "Homograph",
			host: "gооgle.com",
			expected: []string{
				"Homograph domain detected",
			},
		},
		{
			name:     "Normal",
			host:     "google.com",
			expected: []string{},
		},
	}

	for _, tc := range tests {

		t.Run(tc.name, func(t *testing.T) {

			result := DetectHomographDomain(tc.host)

			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("expected %v got %v", tc.expected, result)
			}
		})
	}
}

func TestCheckRedirectURL(t *testing.T) {

	tests := []struct {
		name     string
		url      string
		expected []string
	}{
		{
			name: "Bitly",
			url:  "https://bit.ly/abc123",
			expected: []string{
				"Suspicious URL redirection detected",
			},
		},
		{
			name: "TinyURL",
			url:  "https://tinyurl.com/test",
			expected: []string{
				"Suspicious URL redirection detected",
			},
		},
		{
			name:     "Normal",
			url:      "https://google.com",
			expected: []string{},
		},
	}

	for _, tc := range tests {

		t.Run(tc.name, func(t *testing.T) {

			result := CheckRedirectURL(tc.url)

			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("expected %v got %v", tc.expected, result)
			}
		})
	}
}
package sandbox

import (
	"reflect"
	"testing"
)

func TestExtractSandboxURLs(t *testing.T) {

	tests := []struct {
		name     string
		content  string
		expected []string
	}{
		{
			name:     "Single URL",
			content:  "Visit https://evil.com now",
			expected: []string{"https://evil.com"},
		},
		{
			name:     "Multiple URLs",
			content:  "https://evil.com http://abc.com",
			expected: []string{"https://evil.com", "http://abc.com"},
		},
		{
			name:     "No URL",
			content:  "hello world",
			expected: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			result := ExtractSandboxURLs(tt.content)

			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("expected %v got %v", tt.expected, result)
			}
		})
	}
}

func TestExtractSandboxIPs(t *testing.T) {

	tests := []struct {
		name     string
		content  string
		expected []string
	}{
		{
			name:     "Single IP",
			content:  "8.8.8.8",
			expected: []string{"8.8.8.8"},
		},
		{
			name:     "Multiple IPs",
			content:  "1.1.1.1 8.8.8.8",
			expected: []string{"1.1.1.1", "8.8.8.8"},
		},
		{
			name:     "No IP",
			content:  "hello",
			expected: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			result := ExtractSandboxIPs(tt.content)

			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("expected %v got %v", tt.expected, result)
			}
		})
	}
}

func TestExtractSandboxEmails(t *testing.T) {

	tests := []struct {
		name     string
		content  string
		expected []string
	}{
		{
			name:     "Single Email",
			content:  "abc@test.com",
			expected: []string{"abc@test.com"},
		},
		{
			name:     "Multiple Emails",
			content:  "a@test.com b@test.com",
			expected: []string{"a@test.com", "b@test.com"},
		},
		{
			name:     "No Email",
			content:  "hello",
			expected: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			result := ExtractSandboxEmails(tt.content)

			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("expected %v got %v", tt.expected, result)
			}
		})
	}
}

func TestExtractSandboxDomains(t *testing.T) {

	tests := []struct {
		name     string
		content  string
		expected []string
	}{
		{
			name:     "Single Domain",
			content:  "evil.com",
			expected: []string{"evil.com"},
		},
		{
			name:     "Multiple Domains",
			content:  "evil.com google.com",
			expected: []string{"evil.com", "google.com"},
		},
		{
			name:     "Ignore File Extensions",
			content:  "invoice.exe payload.ps1 test.docm",
			expected: nil,
		},
		{
			name:     "Duplicate Domains",
			content:  "evil.com evil.com",
			expected: []string{"evil.com"},
		},
		{
			name:     "No Domain",
			content:  "hello world",
			expected: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			result := ExtractSandboxDomains(tt.content)

			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("expected %v got %v", tt.expected, result)
			}
		})
	}
}
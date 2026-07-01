package pdfanalyzer

import (
	"reflect"
	"testing"
)

func TestAnalyzePDFText(t *testing.T) {

	tests := []struct {
		name     string
		text     string
		expected []string
	}{
		{
			name: "Login",
			text: "Please login immediately.",
			expected: []string{
				"PDF phishing keyword detected: login",
			},
		},
		{
			name: "Verify Account",
			text: "Verify Account now.",
			expected: []string{
				"PDF phishing keyword detected: verify account",
			},
		},
		{
			name: "Click Here",
			text: "Click here to continue.",
			expected: []string{
				"PDF phishing keyword detected: click here",
			},
		},
		{
			name: "Password",
			text: "Enter your password.",
			expected: []string{
				"PDF phishing keyword detected: password",
			},
		},
		{
			name: "Bank Account",
			text: "Update your bank account.",
			expected: []string{
				"PDF phishing keyword detected: bank account",
			},
		},
		{
			name: "Security Alert",
			text: "Security Alert detected.",
			expected: []string{
				"PDF phishing keyword detected: security alert",
			},
		},
		{
			name: "Multiple Keywords",
			text: "Login now. Click here. Verify account.",
			expected: []string{
				"PDF phishing keyword detected: login",
				"PDF phishing keyword detected: verify account",
				"PDF phishing keyword detected: click here",
			},
		},
		{
			name:     "Safe PDF",
			text:     "Weekly meeting notes.",
			expected: []string{},
		},
		{
			name:     "Empty",
			text:     "",
			expected: []string{},
		},
	}

	for _, tc := range tests {

		t.Run(tc.name, func(t *testing.T) {

			result := AnalyzePDFText(tc.text)

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
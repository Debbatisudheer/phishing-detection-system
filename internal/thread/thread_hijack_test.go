package thread

import (
	"reflect"
	"testing"
)

func TestDetectThreadHijack(t *testing.T) {

	tests := []struct {
		name     string
		subject  string
		body     string
		expected []string
	}{
		{
			name:    "Reply Email",
			subject: "Re: Invoice",
			body:    "",
			expected: []string{
				"Potential thread hijacking detected",
			},
		},
		{
			name:    "Forward Email",
			subject: "Fw: Payment",
			body:    "",
			expected: []string{
				"Potential thread hijacking detected",
			},
		},
		{
			name:    "Forward Email FWD",
			subject: "FWD: Payment",
			body:    "",
			expected: []string{
				"Potential thread hijacking detected",
			},
		},
		{
			name:    "Reply With Keyword",
			subject: "Re: Invoice",
			body:    "Please change bank account immediately.",
			expected: []string{
				"Potential thread hijacking detected",
				"Thread hijack indicator: change bank account",
			},
		},
		{
			name:    "Multiple Indicators",
			subject: "Meeting",
			body:    "Wire transfer and urgent payment with invoice attached.",
			expected: []string{
				"Thread hijack indicator: wire transfer",
				"Thread hijack indicator: urgent payment",
				"Thread hijack indicator: invoice attached",
			},
		},
		{
			name:     "Normal Email",
			subject:  "Project Update",
			body:     "Let's meet tomorrow.",
			expected: []string{},
		},
		{
			name:     "Empty Subject And Body",
			subject:  "",
			body:     "",
			expected: []string{},
		},
	}

	for _, tc := range tests {

		t.Run(tc.name, func(t *testing.T) {

			result := DetectThreadHijack(
				tc.subject,
				tc.body,
			)

			if !reflect.DeepEqual(
				result,
				tc.expected,
			) {
				t.Errorf(
					"expected %v got %v",
					tc.expected,
					result,
				)
			}
		})
	}
}
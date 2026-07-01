package bec

import (
	"reflect"
	"testing"
)

func TestDetectBEC(t *testing.T) {

	tests := []struct {
		name     string
		subject  string
		body     string
		expected []string
	}{
		{
			name:    "Wire Transfer",
			subject: "Wire Transfer Required",
			body:    "",
			expected: []string{
				"BEC indicator detected: wire transfer",
			},
		},
		{
			name:    "Gift Card",
			subject: "",
			body:    "Please purchase a gift card immediately.",
			expected: []string{
				"BEC indicator detected: gift card",
			},
		},
		{
			name:    "Multiple Indicators",
			subject: "Urgent Payment",
			body:    "Please send funds today by bank transfer.",
			expected: []string{
				"BEC indicator detected: bank transfer",
				"BEC indicator detected: urgent payment",
				"BEC indicator detected: send funds",
			},
		},
		{
			name:     "No Indicators",
			subject:  "Meeting Reminder",
			body:     "Let's discuss the project tomorrow.",
			expected: []string{},
		},
		{
			name:    "Case Insensitive",
			subject: "WIRE TRANSFER",
			body:    "",
			expected: []string{
				"BEC indicator detected: wire transfer",
			},
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

			result := DetectBEC(
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
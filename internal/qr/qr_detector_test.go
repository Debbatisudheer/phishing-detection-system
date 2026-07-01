package qr

import (
	"reflect"
	"testing"
)

func TestDetectQRPhishing(t *testing.T) {

	tests := []struct {
		name     string
		subject  string
		body     string
		expected []string
	}{
		{
			name:    "QR Code Subject",
			subject: "QR Code Verification",
			body:    "",
			expected: []string{
				"QR code phishing detected",
			},
		},
		{
			name:    "Scan QR",
			subject: "",
			body:    "Please scan QR immediately.",
			expected: []string{
				"QR code phishing detected",
			},
		},
		{
			name:    "Scan Code",
			subject: "",
			body:    "Scan code to login.",
			expected: []string{
				"QR code phishing detected",
			},
		},
		{
			name:    "Case Insensitive",
			subject: "SCAN QR",
			body:    "",
			expected: []string{
				"QR code phishing detected",
			},
		},
		{
			name:     "Normal Email",
			subject:  "Meeting",
			body:     "Let's meet tomorrow.",
			expected: []string{},
		},
		{
			name:     "Empty",
			subject:  "",
			body:     "",
			expected: []string{},
		},
	}

	for _, tc := range tests {

		t.Run(tc.name, func(t *testing.T) {

			result := DetectQRPhishing(
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
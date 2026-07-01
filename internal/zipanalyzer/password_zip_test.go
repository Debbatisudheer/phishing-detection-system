package zipanalyzer

import (
	"reflect"
	"testing"
)

func TestDetectPasswordProtectedZIP(t *testing.T) {

	tests := []struct {
		name     string
		subject  string
		body     string
		expected []string
	}{
		{
			name:    "Password ZIP",
			subject: "Password for ZIP",
			body:    "",
			expected: []string{
				"Password-protected ZIP suspected",
			},
		},
		{
			name:    "ZIP Password in Body",
			subject: "",
			body:    "The zip password is 1234",
			expected: []string{
				"Password-protected ZIP suspected",
			},
		},
		{
			name:     "Normal Email",
			subject:  "Meeting",
			body:     "See you tomorrow",
			expected: []string{},
		},
	}

	for _, tc := range tests {

		t.Run(tc.name, func(t *testing.T) {

			result := DetectPasswordProtectedZIP(
				tc.subject,
				tc.body,
			)

			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("expected %v got %v", tc.expected, result)
			}
		})
	}
}
package emailauth

import (
	"reflect"
	"testing"
)

func TestCheckEmailAuthentication(t *testing.T) {

	tests := []struct {
		name     string
		sender   string
		expected []string
	}{
		{
			name:   "Gmail Sender",
			sender: "user@gmail.com",
			expected: []string{
				"SPF PASS",
				"DKIM PASS",
				"DMARC PASS",
			},
		},
		{
			name:   "Mixed Case Gmail",
			sender: "User@GMAIL.COM",
			expected: []string{
				"SPF PASS",
				"DKIM PASS",
				"DMARC PASS",
			},
		},
		{
			name:   "Corporate Domain",
			sender: "user@company.com",
			expected: []string{
				"SPF FAIL",
				"DKIM FAIL",
				"DMARC FAIL",
			},
		},
		{
			name:   "Empty Sender",
			sender: "",
			expected: []string{
				"SPF FAIL",
				"DKIM FAIL",
				"DMARC FAIL",
			},
		},
		{
			name:   "Yahoo Sender",
			sender: "user@yahoo.com",
			expected: []string{
				"SPF FAIL",
				"DKIM FAIL",
				"DMARC FAIL",
			},
		},
	}

	for _, tc := range tests {

		t.Run(tc.name, func(t *testing.T) {

			result := CheckEmailAuthentication(tc.sender)

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
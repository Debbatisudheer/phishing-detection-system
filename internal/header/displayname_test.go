package header

import (
	"reflect"
	"testing"
)

func TestDetectDisplayNameSpoofing(t *testing.T) {

	tests := []struct {
		name     string
		from     string
		expected []string
	}{
		{
			name: "Microsoft Spoof",
			from: "Microsoft Support <attacker@gmail.com>",
			expected: []string{
				"Display name spoofing detected: microsoft",
			},
		},
		{
			name: "PayPal Spoof",
			from: "PayPal Team <evil@yahoo.com>",
			expected: []string{
				"Display name spoofing detected: paypal",
			},
		},
		{
			name:     "Legitimate Microsoft",
			from:     "admin@microsoft.com",
			expected: []string{},
		},
		{
			name:     "No Brand",
			from:     "alice@example.com",
			expected: []string{},
		},
	}

	for _, tc := range tests {

		t.Run(tc.name, func(t *testing.T) {

			result := DetectDisplayNameSpoofing(
				tc.from,
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
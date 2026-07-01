package sender

import (
	"reflect"
	"testing"
)

func TestCheckSenderReputation(t *testing.T) {

	tests := []struct {
		name     string
		sender   string
		expected []string
	}{
		{
			name:   "Evil Sender",
			sender: "evil@example.com",
			expected: []string{
				"Suspicious sender reputation detected",
			},
		},
		{
			name:   "Phishing Sender",
			sender: "phishing@test.com",
			expected: []string{
				"Suspicious sender reputation detected",
			},
		},
		{
			name:   "Attacker Sender",
			sender: "attacker@abc.com",
			expected: []string{
				"Suspicious sender reputation detected",
			},
		},
		{
			name:   "Mixed Case",
			sender: "EvIl@Example.com",
			expected: []string{
				"Suspicious sender reputation detected",
			},
		},
		{
			name:     "Legitimate Sender",
			sender:   "alice@example.com",
			expected: []string{},
		},
	}

	for _, tc := range tests {

		t.Run(tc.name, func(t *testing.T) {

			result := CheckSenderReputation(
				tc.sender,
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
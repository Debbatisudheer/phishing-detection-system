package decision

import "testing"

func TestMakeDecision(t *testing.T) {

	tests := []struct {
		name     string
		score    int
		expected string
	}{
		{
			name:     "Allow Zero",
			score:    0,
			expected: "ALLOW",
		},
		{
			name:     "Allow Ninety Nine",
			score:    99,
			expected: "ALLOW",
		},
		{
			name:     "Suspicious Boundary",
			score:    100,
			expected: "SUSPICIOUS",
		},
		{
			name:     "Suspicious Mid",
			score:    250,
			expected: "SUSPICIOUS",
		},
		{
			name:     "Suspicious Upper Boundary",
			score:    399,
			expected: "SUSPICIOUS",
		},
		{
			name:     "Quarantine Boundary",
			score:    400,
			expected: "QUARANTINE",
		},
		{
			name:     "Quarantine High",
			score:    1000,
			expected: "QUARANTINE",
		},
	}

	for _, tc := range tests {

		t.Run(tc.name, func(t *testing.T) {

			result := MakeDecision(tc.score)

			if result != tc.expected {
				t.Errorf(
					"expected %s got %s",
					tc.expected,
					result,
				)
			}
		})
	}
}
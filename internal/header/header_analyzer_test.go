package header

import (
	"reflect"
	"testing"
)

func TestAnalyzeHeaders(t *testing.T) {

	tests := []struct {
		name       string
		from       string
		replyTo    string
		returnPath string
		expected   []string
	}{
		{
			name:       "No Mismatch",
			from:       "alice@example.com",
			replyTo:    "alice@example.com",
			returnPath: "alice@example.com",
			expected:   []string{},
		},
		{
			name:       "Reply-To Mismatch",
			from:       "alice@example.com",
			replyTo:    "attacker@example.com",
			returnPath: "alice@example.com",
			expected: []string{
				"Reply-To mismatch detected",
			},
		},
		{
			name:       "Return-Path Mismatch",
			from:       "alice@example.com",
			replyTo:    "alice@example.com",
			returnPath: "bounce@example.com",
			expected: []string{
				"Return-Path mismatch detected",
			},
		},
		{
			name:       "Both Mismatches",
			from:       "alice@example.com",
			replyTo:    "attacker@example.com",
			returnPath: "bounce@example.com",
			expected: []string{
				"Reply-To mismatch detected",
				"Return-Path mismatch detected",
			},
		},
	}

	for _, tc := range tests {

		t.Run(tc.name, func(t *testing.T) {

			result := AnalyzeHeaders(
				tc.from,
				tc.replyTo,
				tc.returnPath,
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
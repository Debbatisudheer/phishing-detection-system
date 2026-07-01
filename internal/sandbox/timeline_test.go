package sandbox

import (
	"strings"
	"testing"
)

func TestBuildTimeline(t *testing.T) {

	tests := []struct {
		name      string
		findings  []string
		expectLen int
	}{
		{
			name: "Single Event",
			findings: []string{
				"PowerShell execution",
			},
			expectLen: 1,
		},
		{
			name: "Multiple Events",
			findings: []string{
				"PowerShell execution",
				"Network communication",
				"Dropped file",
			},
			expectLen: 3,
		},
		{
			name:      "Empty",
			findings:  []string{},
			expectLen: 0,
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {

			result := BuildTimeline(tt.findings)

			if len(result) != tt.expectLen {

				t.Fatalf(
					"expected %d events got %d",
					tt.expectLen,
					len(result),
				)
			}

			for _, event := range result {

				if !strings.HasPrefix(
					event,
					"[00:",
				) {

					t.Errorf(
						"timeline format incorrect: %s",
						event,
					)
				}
			}
		})
	}
}
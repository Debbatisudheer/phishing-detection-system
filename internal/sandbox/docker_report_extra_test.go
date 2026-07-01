package sandbox

import (
	"errors"
	"testing"
)

func TestBuildDockerReportBranches(t *testing.T) {

	tests := []struct {
		name     string
		err      error
		expected int
	}{
		{
			name:     "Success",
			err:      nil,
			expected: 4,
		},
		{
			name:     "Failure",
			err:      errors.New("docker failed"),
			expected: 4,
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {

			report := BuildDockerReport(
				"dummy output",
				tt.err,
			)

			if len(report) != tt.expected {
				t.Fatalf(
					"expected %d findings got %d",
					tt.expected,
					len(report),
				)
			}
		})
	}
}
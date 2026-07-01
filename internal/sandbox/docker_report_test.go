package sandbox

import (
	"errors"
	"reflect"
	"testing"
)

func TestBuildDockerReport(t *testing.T) {

	tests := []struct {
		name     string
		err      error
		expected []string
	}{
		{
			name: "Successful Execution",
			err:  nil,
			expected: []string{
				"Docker Event: Container Started",
				"Docker Event: File Mounted",
				"Docker Event: Execution Successful",
				"Docker Event: Container Destroyed",
			},
		},
		{
			name: "Failed Execution",
			err:  errors.New("docker failed"),
			expected: []string{
				"Docker Event: Container Started",
				"Docker Event: File Mounted",
				"Docker Event: Execution Failed",
				"Docker Event: Container Destroyed",
			},
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {

			result := BuildDockerReport(
				"",
				tt.err,
			)

			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf(
					"expected %v got %v",
					tt.expected,
					result,
				)
			}
		})
	}
}
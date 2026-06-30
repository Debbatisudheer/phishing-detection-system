package parser

import (
	"reflect"
	"testing"
)

func TestExtractAttachmentNames(t *testing.T) {

	tests := []struct {
		name        string
		contentType string
		expected    []string
	}{
		{
			name:        "Valid Filename",
			contentType: `attachment; filename="invoice.pdf"`,
			expected: []string{
				"invoice.pdf",
			},
		},
		{
			name:        "No Filename",
			contentType: `attachment`,
			expected:    []string{},
		},
		{
			name:        "Invalid Content Type",
			contentType: `@@@@@@`,
			expected:    []string{},
		},
	}

	for _, tc := range tests {

		t.Run(tc.name, func(t *testing.T) {

			result := ExtractAttachmentNames(
				tc.contentType,
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
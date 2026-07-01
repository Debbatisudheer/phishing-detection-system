package urlanalyzer

import (
	"reflect"
	"testing"
)

func TestDetectShortenedURL(t *testing.T) {

	tests := []struct {
		name     string
		url      string
		expected []string
	}{
		{
			name: "Bitly",
			url:  "https://bit.ly/abc123",
			expected: []string{
				"Shortened URL detected: bit.ly",
			},
		},
		{
			name: "TinyURL",
			url:  "https://tinyurl.com/test",
			expected: []string{
				"Shortened URL detected: tinyurl.com",
			},
		},
		{
			name: "TCO",
			url:  "https://t.co/xyz",
			expected: []string{
				"Shortened URL detected: t.co",
			},
		},
		{
			name: "GOOGL",
			url:  "https://goo.gl/test",
			expected: []string{
				"Shortened URL detected: goo.gl",
			},
		},
		{
			name: "ISGD",
			url:  "https://is.gd/test",
			expected: []string{
				"Shortened URL detected: is.gd",
			},
		},
		{
			name: "Case Insensitive",
			url:  "HTTPS://BIT.LY/ABC",
			expected: []string{
				"Shortened URL detected: bit.ly",
			},
		},
		{
			name:     "Normal URL",
			url:      "https://google.com",
			expected: []string{},
		},
		{
			name:     "Empty URL",
			url:      "",
			expected: []string{},
		},
	}

	for _, tc := range tests {

		t.Run(tc.name, func(t *testing.T) {

			result := DetectShortenedURL(tc.url)

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
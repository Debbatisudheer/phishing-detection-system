package phishtank

import (
	"os"
	"reflect"
	"testing"
)

func TestLoadPhishTankFeed(t *testing.T) {

	t.Run("Valid Feed", func(t *testing.T) {

		tmp, err := os.CreateTemp("", "phishtank-*.txt")
		if err != nil {
			t.Fatal(err)
		}
		defer os.Remove(tmp.Name())

		tmp.WriteString("evil.com\n")
		tmp.WriteString("Phishing.com\n")
		tmp.Close()

		result := LoadPhishTankFeed(tmp.Name())

		expected := []string{
			"evil.com",
			"phishing.com",
		}

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("expected %v got %v", expected, result)
		}
	})

	t.Run("Missing File", func(t *testing.T) {

		result := LoadPhishTankFeed("does-not-exist.txt")

		if len(result) != 0 {
			t.Errorf("expected empty slice got %v", result)
		}
	})
}

func TestCheckPhishTank(t *testing.T) {

	// Override the global feed for testing
	PhishTankDomains = []string{
		"evil.com",
		"phishing.com",
	}

	tests := []struct {
		name     string
		url      string
		expected []string
	}{
		{
			name: "Known Domain",
			url:  "https://evil.com/login",
			expected: []string{
				"PhishTank hit: evil.com",
			},
		},
		{
			name: "Second Domain",
			url:  "https://phishing.com/index.html",
			expected: []string{
				"PhishTank hit: phishing.com",
			},
		},
		{
			name: "Case Insensitive",
			url:  "HTTPS://EVIL.COM/LOGIN",
			expected: []string{
				"PhishTank hit: evil.com",
			},
		},
		{
			name:     "Safe Domain",
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

			result := CheckPhishTank(tc.url)

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
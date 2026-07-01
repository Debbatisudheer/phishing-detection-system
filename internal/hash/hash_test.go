package hash

import (
	"os"
	"reflect"
	"testing"
)

func TestCalculateSHA256(t *testing.T) {

	t.Run("Valid File", func(t *testing.T) {

		tmp, err := os.CreateTemp("", "hash-test-*.txt")
		if err != nil {
			t.Fatal(err)
		}
		defer os.Remove(tmp.Name())

		_, err = tmp.WriteString("hello world")
		if err != nil {
			t.Fatal(err)
		}

		tmp.Close()

		hash := CalculateSHA256(tmp.Name())

		if hash == "" {
			t.Error("expected non-empty hash")
		}

		if len(hash) != 64 {
			t.Errorf("expected hash length 64 got %d", len(hash))
		}
	})

	t.Run("Missing File", func(t *testing.T) {

		hash := CalculateSHA256("file-does-not-exist.txt")

		if hash != "" {
			t.Errorf("expected empty string got %s", hash)
		}
	})
}

func TestCheckHashReputation(t *testing.T) {

	tests := []struct {
		name     string
		hash     string
		expected []string
	}{
		{
			name: "Known Malware Hash",
			hash: "091b572abf984382a95b9455e407f284e24bfb616890abc3fcdb2be68813f39d",
			expected: []string{
				"Known malware hash detected",
			},
		},
		{
			name:     "Unknown Hash",
			hash:     "abcdef123456",
			expected: []string{},
		},
		{
			name:     "Empty Hash",
			hash:     "",
			expected: []string{},
		},
	}

	for _, tc := range tests {

		t.Run(tc.name, func(t *testing.T) {

			result := CheckHashReputation(tc.hash)

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
package threatfeed

import (
	"os"
	"reflect"
	"testing"
)

func TestLoadFeed(t *testing.T) {

	t.Run("Valid Feed", func(t *testing.T) {

		tmp, err := os.CreateTemp("", "feed-*.txt")
		if err != nil {
			t.Fatal(err)
		}
		defer os.Remove(tmp.Name())

		tmp.WriteString("evil.com\n")
		tmp.WriteString("phishing.com\n")
		tmp.Close()

		result := LoadFeed(tmp.Name())

		expected := []string{
			"evil.com",
			"phishing.com",
		}

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("expected %v got %v", expected, result)
		}
	})

	t.Run("Missing File", func(t *testing.T) {

		result := LoadFeed("does-not-exist.txt")

		if len(result) != 0 {
			t.Errorf("expected empty slice got %v", result)
		}
	})
}
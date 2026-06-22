package sandbox

import "strings"

func Contains(
	content string,
	value string,
) bool {

	return strings.Contains(
		strings.ToLower(content),
		strings.ToLower(value),
	)
}
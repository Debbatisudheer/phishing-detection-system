package parser

import (
	"regexp"
)

func ExtractURLs(text string) []string {

	re := regexp.MustCompile(`https?://[^\s]+`)

	urls := re.FindAllString(text, -1)

	return urls
}
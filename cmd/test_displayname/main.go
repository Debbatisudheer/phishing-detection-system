package main

import (
	"fmt"

	"phishing-platform/internal/header"
)

func main() {

	findings :=
		header.DetectDisplayNameSpoofing(
			"Microsoft Security <attacker@gmail.com>",
		)

	fmt.Println(
		findings,
	)
}
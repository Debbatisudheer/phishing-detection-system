package main

import (
	"fmt"

	"phishing-platform/internal/domain"
)

func main() {

	findings :=
		domain.DetectHomographDomain(
			"microsоft.com",
		)

	fmt.Println(
		"Findings:",
		findings,
	)
}
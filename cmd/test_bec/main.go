package main

import (
	"fmt"

	"phishing-platform/internal/bec"
)

func main() {

	findings :=
		bec.DetectBEC(
			"Urgent Payment",
			"Please complete wire transfer today",
		)

	fmt.Println(
		findings,
	)
}
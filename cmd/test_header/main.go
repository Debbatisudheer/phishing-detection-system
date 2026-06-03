package main

import (
	"fmt"

	"phishing-platform/internal/header"
)

func main() {

	findings :=
		header.AnalyzeHeaders(
			"support@microsoft.com",
			"attacker@gmail.com",
			"bounce@evil.com",
		)

	fmt.Println(
		"Header Findings:",
		findings,
	)
}
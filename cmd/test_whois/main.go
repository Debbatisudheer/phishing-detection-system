package main

import (
	"fmt"

	"phishing-platform/internal/whois"
)

func main() {

	findings :=
		whois.AnalyzeDomain(
			"microsoft-security.xyz",
		)

	fmt.Println(
		"WHOIS Findings:",
		findings,
	)
}
package main

import (
	"fmt"

	"phishing-platform/internal/dnsreputation"
)

func main() {

	findings :=
		dnsreputation.CheckDNSReputation(
			"evil.com",
		)

	fmt.Println(
		findings,
	)
}
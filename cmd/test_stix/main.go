package main

import (
	"fmt"

	"phishing-platform/internal/stix"
)

func main() {

	err :=
		stix.ExportURLIndicator(
			"https://evil.com/login",
			"stix_indicator.json",
		)

	if err != nil {
		panic(err)
	}

	fmt.Println(
		"STIX Indicator Generated",
	)
}
package main

import (
	"fmt"

	"phishing-platform/internal/virustotal"
)

func main() {

	response, err :=
		virustotal.QueryURL(
			"https://evil.com/login",
		)

	if err != nil {
		panic(err)
	}

	findings :=
		virustotal.CheckURLReputation(
			response,
		)

	fmt.Println(
		"VT URL Findings:",
		findings,
	)
}
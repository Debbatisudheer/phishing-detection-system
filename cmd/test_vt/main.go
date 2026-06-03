package main

import (
	"fmt"

	"phishing-platform/internal/virustotal"
)

func main() {

	hash :=
		"44d88612fea8a8f36de82e1278abb02f"

	response, err :=
		virustotal.QueryHash(
			hash,
		)

	if err != nil {
		panic(err)
	}

	findings :=
		virustotal.CheckHashReputation(
			response,
		)

	fmt.Println(
		"VT Findings:",
		findings,
	)
}
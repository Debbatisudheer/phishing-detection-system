package main

import (
	"fmt"

	"phishing-platform/internal/sigma"
)

func main() {

	err := sigma.GenerateRule(
		"attacker@evil.com",
		775,
		"attack.t1566",
		"phishing_rule.yml",
	)

	if err != nil {
		panic(err)
	}

	fmt.Println(
		"Sigma Rule Generated",
	)
}
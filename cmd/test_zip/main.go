package main

import (
	"fmt"

	"phishing-platform/internal/risk"
)

func main() {

	findings := []string{
		"Nested ZIP detected: inner.zip",
	}

	score :=
		risk.CalculateRisk(
			"Test",
			"Test",
			nil,
			findings,
		)

	fmt.Println(
		"FINAL TEST SCORE:",
		score,
	)
}
package main

import (
	"fmt"

	"phishing-platform/internal/macroanalyzer"
)

func main() {

	text :=
		macroanalyzer.ExtractMacroText(
			"uploads/invoice.docm",
		)

	fmt.Println(
		"MACRO CONTENT:",
		text,
	)
}
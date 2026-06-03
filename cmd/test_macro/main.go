package main

import (
	"fmt"

	"phishing-platform/internal/macroanalyzer"
)

func main() {

	macroText :=
		macroanalyzer.ExtractWPSMacroText(
			"uploads\\real_invoicee.docm",
		)

	fmt.Println(
		"EXTRACTED MACRO CONTENT:",
	)

	fmt.Println(
		macroText,
	)

	findings :=
		macroanalyzer.AnalyzeMacroContent(
			macroText,
		)

	fmt.Println(
		"FINDINGS:",
		findings,
	)
}
package main

import (
	"fmt"

	"phishing-platform/internal/pdfanalyzer"
)

func main() {

	text := `
Security Alert

Please verify account immediately.

Click here to login.
`

	findings :=
		pdfanalyzer.AnalyzePDFText(
			text,
		)

	fmt.Println(
		"PDF Findings:",
		findings,
	)
}
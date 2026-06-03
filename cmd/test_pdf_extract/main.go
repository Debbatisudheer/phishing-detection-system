package main

import (
	"fmt"

	"phishing-platform/internal/pdfanalyzer"
)

func main() {

	

	text :=
		pdfanalyzer.ExtractPDFText(
			"uploads/invoice.pdf",
		)

	fmt.Println(
		"PDF Text:",
		text,
	)
}
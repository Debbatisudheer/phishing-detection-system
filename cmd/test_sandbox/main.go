package main

import (
	"fmt"

	"phishing-platform/internal/sandbox"
)

func main() {

	content := `
powershell.exe
CreateObject("WScript.Shell")
https://evil.com
`

	findings :=
		sandbox.AnalyzeBehavior(
			content,
		)

	fmt.Println(
		"Sandbox Findings:",
		findings,
	)
}
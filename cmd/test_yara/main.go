package main

import (
	"fmt"

	"phishing-platform/internal/yara"
)

func main() {

	content := `
AutoOpen
CreateObject("WScript.Shell")
powershell.exe
https://evil.com
`

	findings :=
		yara.ScanContent(
			content,
		)

	fmt.Println(
		"YARA Findings:",
		findings,
	)
}
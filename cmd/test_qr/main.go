package main

import (
	"fmt"

	"phishing-platform/internal/qr"
)

func main() {

	results :=
		qr.DecodeQRImage(
			"qrscan.png",
		)

	fmt.Println(
		"QR Results:",
		results,
	)
}
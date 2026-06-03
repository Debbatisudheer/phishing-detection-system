package qr

import "strings"

func DetectQRPhishing(
	subject string,
	body string,
) []string {

	var findings []string

	text :=
		strings.ToLower(
			subject + " " + body,
		)

	qrDetected := false

	if strings.Contains(text, "qr code") {
		qrDetected = true
	}

	if strings.Contains(text, "scan qr") {
		qrDetected = true
	}

	if strings.Contains(text, "scan code") {
		qrDetected = true
	}

	if qrDetected {

		findings = append(
			findings,
			"QR code phishing detected",
		)
	}

	return findings
}
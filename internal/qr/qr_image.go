package qr

import (
	"image"
	"os"

	_ "image/jpeg"
	_ "image/png"

	"github.com/liyue201/goqr"
)

func DecodeQRImage(
	filePath string,
) []string {

	var findings []string

	file, err := os.Open(filePath)
	if err != nil {
		return findings
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return findings
	}

	qrcodes, err := goqr.Recognize(img)
	if err != nil {
		return findings
	}

	for _, qr := range qrcodes {
		findings = append(
			findings,
			string(qr.Payload),
		)
	}

	return findings
}
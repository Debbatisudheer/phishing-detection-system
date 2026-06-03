package pdfanalyzer

import (
	"fmt"

	"github.com/ledongthuc/pdf"
)

func ExtractPDFText(
	filePath string,
) string {

	f, r, err :=
		pdf.Open(
			filePath,
		)

	if err != nil {

		fmt.Println(
			"PDF OPEN ERROR:",
			err,
		)

		return ""
	}

	defer f.Close()

	fmt.Println(
		"Pages:",
		r.NumPage(),
	)

	var text string

	totalPage :=
		r.NumPage()

	for i := 1; i <= totalPage; i++ {

		page :=
			r.Page(i)

		if page.V.IsNull() {
			continue
		}

		pageText, err :=
			page.GetPlainText(
				nil,
			)

		if err == nil {

			text += pageText
		}
	}

	return text
}
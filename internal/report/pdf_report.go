package report

import (
	"fmt"

	"github.com/phpdave11/gofpdf"
)

func GeneratePDFReport(
	fileName string,
	riskScore int,
	riskLevel string,
	verdict string,
	findings string,
	mitre string,
) error {

	pdf := gofpdf.New(
		"P",
		"mm",
		"A4",
		"",
	)

	pdf.AddPage()

	pdf.SetFont(
		"Arial",
		"B",
		16,
	)

	pdf.Cell(
		190,
		10,
		"PHISHING INVESTIGATION REPORT",
	)

	pdf.Ln(15)

	pdf.SetFont(
		"Arial",
		"",
		12,
	)

	pdf.Cell(
		190,
		10,
		fmt.Sprintf(
			"File Name: %s",
			fileName,
		),
	)

	pdf.Ln(8)

	pdf.Cell(
		190,
		10,
		fmt.Sprintf(
			"Risk Score: %d",
			riskScore,
		),
	)

	pdf.Ln(8)

	pdf.Cell(
		190,
		10,
		fmt.Sprintf(
			"Risk Level: %s",
			riskLevel,
		),
	)

	pdf.Ln(8)

	pdf.Cell(
		190,
		10,
		fmt.Sprintf(
			"Verdict: %s",
			verdict,
		),
	)

	pdf.Ln(15)

	pdf.MultiCell(
		190,
		8,
		"MITRE ATT&CK:\n"+mitre,
		"",
		"",
		false,
	)

	pdf.Ln(5)

	pdf.MultiCell(
		190,
		8,
		"Findings:\n"+findings,
		"",
		"",
		false,
	)

	return pdf.OutputFileAndClose(
		"investigation_report.pdf",
	)
}
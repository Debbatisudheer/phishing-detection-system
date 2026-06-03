package main

import (
	"fmt"

	"phishing-platform/internal/ioc"
)

func main() {

	report := ioc.IOCReport{
		Sender: "attacker@evil.com",
		URLs: []string{
			"https://evil.com/login",
		},
		Hashes: []string{
			"091b572abf984382",
		},
		Attachments: []string{
			"invoice.zip",
		},
		MITRE: "T1566 - Phishing",
	}

	err :=
		ioc.ExportIOC(
			report,
			"ioc_report.json",
		)

	if err != nil {
		panic(err)
	}

	fmt.Println(
		"IOC Report Generated",
	)
}
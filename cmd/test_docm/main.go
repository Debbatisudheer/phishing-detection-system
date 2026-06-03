package main

import (
	"archive/zip"
	"fmt"
)

func main() {

	reader, err :=
	zip.OpenReader(
		"uploads\\real_invoicee.docm",
	)

	if err != nil {

		fmt.Println(
			"ZIP OPEN ERROR:",
			err,
		)

		return
	}

	defer reader.Close()

	fmt.Println(
		"DOCM FILES:",
	)

	for _, file := range reader.File {

	fmt.Println(
		file.Name,
	)

	if file.Name == "word/vbaProject.bin" {

		fmt.Println(
			"VBA PROJECT FOUND",
		)
	}
}


}
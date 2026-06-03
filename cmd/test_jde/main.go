package main

import (
	"archive/zip"
	"fmt"
	"io"
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

	for _, file := range reader.File {

		if file.Name ==
			"word/JDEData.bin" {

			fmt.Println(
				"FOUND JDEData.bin",
			)

			rc, _ :=
				file.Open()

			data, _ :=
				io.ReadAll(
					rc,
				)

			rc.Close()

			fmt.Println(
				string(data),
			)
		}
	}
}
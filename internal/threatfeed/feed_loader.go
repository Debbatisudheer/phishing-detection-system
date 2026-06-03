package threatfeed

import (
	"bufio"
	"os"
)

func LoadFeed(
	filePath string,
) []string {

	var domains []string

	file, err :=
		os.Open(
			filePath,
		)

	if err != nil {
		return domains
	}

	defer file.Close()

	scanner :=
		bufio.NewScanner(
			file,
		)

	for scanner.Scan() {

		domains = append(
			domains,
			scanner.Text(),
		)
	}

	return domains
}
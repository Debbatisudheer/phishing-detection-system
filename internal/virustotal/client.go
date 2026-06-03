package virustotal

import (
	"io"
	"net/http"
	"os"
)

func QueryHash(
	hash string,
) ([]byte, error) {

	apiKey :=
		os.Getenv(
			"VT_API_KEY",
		)

	url :=
		"https://www.virustotal.com/api/v3/files/" +
			hash

	req, err :=
		http.NewRequest(
			"GET",
			url,
			nil,
		)

	if err != nil {
		return nil, err
	}

	req.Header.Set(
		"x-apikey",
		apiKey,
	)

	client := &http.Client{}

	resp, err :=
		client.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	return io.ReadAll(
		resp.Body,
	)
}